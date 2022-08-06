package heap

import (
	"fmt"
	"go-jvm/src/ch08/classfile"
	"go-jvm/src/ch08/classpath"
	"go-jvm/src/ch08/classpath/entry"
)

type ClassLoader struct {
	classpath *classpath.Classpath
	classMap  map[string]*Class
	verbose   bool
}

func NewClassLoader(classpath *classpath.Classpath, verbose bool) *ClassLoader {
	return &ClassLoader{
		classpath: classpath,
		classMap:  make(map[string]*Class),
		verbose:   verbose,
	}
}

func (c *ClassLoader) LoadClass(name string) *Class {
	if class, ok := c.classMap[name]; ok {
		return class
	}
	if name[0] == '[' {
		return c.loadArrayClass(name)
	}
	return c.loadNonArrayClass(name)
}

func (c *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, //ToDo
		name:        name,
		classLoader: c,
		initStarted: true,
		superClass:  c.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			c.LoadClass("java/lang/Cloneable"),
			c.LoadClass("java/io/Serializable"),
		},
	}
	c.classMap[name] = class
	return class
}

func (c *ClassLoader) loadNonArrayClass(name string) *Class {
	data, ent := c.readClass(name)
	class := c.defineClass(data)
	link(class)
	if c.verbose {
		fmt.Printf("[Loaded %s from %s]\n", name, ent)
	}
	return class
}

func (c *ClassLoader) readClass(name string) ([]byte, entry.Entry) {
	data, ent, err := c.classpath.ReadClass(name)
	if err != nil {
		panic(fmt.Sprintf("java.lang.ClassNotFoundException: %s", name))
	}
	return data, ent
}

func (c *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.classLoader = c
	resolveSuperClass(class)
	resolveInterfaces(class)
	c.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.classLoader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.classLoader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//ToDo: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.10.1
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	rtcp := class.rtcp
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.descriptor {
		case "Z", "B", "C", "S", "I":
			val := rtcp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := rtcp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := rtcp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := rtcp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := rtcp.GetConstant(cpIndex).(string)
			jvStr := JvString(class.classLoader, goStr)
			vars.SetRef(slotId, jvStr)
		}
	}
}
