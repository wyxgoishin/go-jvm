package heap

import (
	"go-jvm/src/ch08/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	rtcp              *RuntimeConstantPool
	fields            []*Field
	methods           []*Method
	classLoader       *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
}

func newClass(cf *classfile.ClassFile) *Class {
	class := new(Class)
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.rtcp = newRuntimeConstantPool(class, cf.ConstantPool())
	class.fields = newField(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (c *Class) InitStarted() bool {
	return c.initStarted
}

func (c *Class) StartInit() {
	c.initStarted = true
}

func (c *Class) GetPackageName() string {
	if idx := strings.LastIndex(c.name, "/"); idx >= 0 {
		return c.name[:idx]
	}
	return ""
}

func (c *Class) isAccessibleTo(other *Class) bool {
	return c.IsPublic() || c.GetPackageName() == other.GetPackageName()
}

func (c *Class) IsSubClassOf(other *Class) bool {
	this := c
	for this != nil {
		if this.superClass == other {
			return true
		}
		this = this.superClass
	}
	return false
}

func (c *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(c)
}

func (c *Class) AccessFlags() uint16 {
	return c.accessFlags
}

func (c *Class) NewObject() *Object {
	return newObject(c)
}

func (c *Class) RuntimeConstantPool() *RuntimeConstantPool {
	return c.rtcp
}

func (c *Class) StaticVars() Slots {
	return c.staticVars
}

// other cast to this
func (c *Class) isAssignableFrom(other *Class) bool {
	sub, this := other, c
	if sub == this {
		return true
	}
	if !sub.IsArray() {
		if !sub.IsInterface() {
			if !this.IsInterface() {
				return sub.IsSubClassOf(this)
			} else {
				return sub.IsImplements(this)
			}
		} else {
			if !this.IsInterface() {
				return this.isJvObject()
			} else {
				return this.isSuperInterfaceOf(sub)
			}
		}
	} else {
		if !this.IsArray() {
			// array can be cast to Object or Cloneable or Serializable
			if !this.IsInterface() {
				return this.isJvObject()
			} else {
				return this.isJvCloneable() || this.isJvSerializable()
			}
		} else {
			sElemClass := sub.ComponentClass()
			thisElemClass := this.ComponentClass()
			return sElemClass == thisElemClass || thisElemClass.isAssignableFrom(sElemClass)
		}
	}
}

func (c *Class) isJvObject() bool {
	return c.name == "java/lang/Object"
}

func (c *Class) isJvCloneable() bool {
	return c.name == "java/lang/Cloneable"
}

func (c *Class) isJvSerializable() bool {
	return c.name == "java/io/Serializable"
}

func (c *Class) IsImplements(other *Class) bool {
	this := c
	for this != nil {
		for _, intface := range this.interfaces {
			if intface == other || intface.isSubInterfaceOf(other) {
				return true
			}
		}
		this = this.superClass
	}
	return false
}

func (c *Class) isSubInterfaceOf(other *Class) bool {
	for _, intface := range c.interfaces {
		if intface == other || intface.isSubInterfaceOf(other) {
			return true
		}
	}
	return false
}

func (c *Class) isSuperInterfaceOf(other *Class) bool {
	return other.isSubInterfaceOf(c)
}

func (c *Class) GetMainMethod() *Method {
	return c.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (c *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range c.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (c *Class) SuperClass() *Class {
	return c.superClass
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) ClassLoader() *ClassLoader {
	return c.classLoader
}

func (c *Class) GetClinitMethod() *Method {
	return c.getStaticMethod("<clinit>", "()V")
}

func (c *Class) IsPublic() bool {
	return c.accessFlags&ACC_PUBLIC > 0
}

func (c *Class) IsFinal() bool {
	return c.accessFlags&ACC_FINAL > 0
}

func (c *Class) IsSuper() bool {
	return c.accessFlags&ACC_SUPER > 0
}

func (c *Class) IsInterface() bool {
	return c.accessFlags&ACC_INTERFACE > 0
}

func (c *Class) IsAbstract() bool {
	return c.accessFlags&ACC_ABSTRACT > 0
}

func (c *Class) IsSynthetic() bool {
	return c.accessFlags&ACC_SYNTHETIC > 0
}

func (c *Class) IsAnnotation() bool {
	return c.accessFlags&ACC_ANNOTATION > 0
}

func (c *Class) IsEnum() bool {
	return c.accessFlags&ACC_ENUM > 0
}

func (c *Class) IsProtected() bool {
	return c.accessFlags&ACC_PROTECTED > 0
}

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"char":    "C",
	"short":   "S",
	"int":     "I",
	"float":   "F",
	"long":    "L",
	"double":  "D",
}

func (c *Class) IsArray() bool {
	return c.name[0] == '['
}

// getElemName, for example, int[] -> [I, String[] -> [Ljava/lang/String;
func (c *Class) ArrayClass() *Class {
	arrClassName := getArrayClassName(c.name)
	return c.classLoader.LoadClass(arrClassName)
}

func getArrayClassName(name string) string {
	return "[" + toDescriptor(name)
}

func toDescriptor(name string) string {
	if name[0] == '[' {
		return name
	}
	if d, ok := primitiveTypes[name]; ok {
		return d
	}
	return "L" + name + ";"
}

func (c *Class) ComponentClass() *Class {
	componetClassName := getComponentClassName(c.name)
	return c.classLoader.LoadClass(componetClassName)
}

func getComponentClassName(name string) string {
	if name[0] == '[' {
		componentTypeDescriptor := name[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array type: " + name)
}

func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}
	for name, d := range primitiveTypes {
		if d == descriptor {
			return name
		}
	}
	panic("Unexpected descriptor: " + descriptor)
}

func (c *Class) getField(name, descriptor string, isStatic bool) *Field {
	for this := c; this != nil; this = this.superClass {
		for _, field := range this.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}
