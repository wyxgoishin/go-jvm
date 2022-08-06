package heap

import (
	"go-jvm/src/ch06/classfile"
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
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
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

func (c *Class) getPackageName() string {
	if idx := strings.LastIndex(c.name, "/"); idx >= 0 {
		return c.name[:idx]
	}
	return ""
}

func (c *Class) isAccessibleTo(other *Class) bool {
	return IsPublic(c.accessFlags) || c.getPackageName() == other.getPackageName()
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

func (c *Class) isAssignableFrom(other *Class) bool {
	s, t := other, c
	if s == t {
		return true
	}
	if !IsInterface(c.accessFlags) {
		return s.IsSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}

func (c *Class) isImplements(other *Class) bool {
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

func (c *Class) GetMainMethod() *Method {
	return c.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (c *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range c.methods {
		if IsStatic(method.accessFlags) && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}
