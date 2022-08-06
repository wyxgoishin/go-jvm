package heap

import (
	"go-jvm/src/ch11/classfile"
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
	jvClass           *Object // java.lang.Class Instance
	sourceFile        string  // 不是每个字节码文件都有该属性
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
	class.sourceFile = getSourceFile(cf)
	return class
}

func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown"
}

func (cls *Class) JvClass() *Object {
	return cls.jvClass
}

func (cls *Class) JvName() string {
	return strings.Replace(cls.name, "/", ".", -1)
}

func (cls *Class) InitStarted() bool {
	return cls.initStarted
}

func (cls *Class) StartInit() {
	cls.initStarted = true
}

func (cls *Class) GetPackageName() string {
	if idx := strings.LastIndex(cls.name, "/"); idx >= 0 {
		return cls.name[:idx]
	}
	return ""
}

func (cls *Class) isAccessibleTo(other *Class) bool {
	return cls.IsPublic() || cls.GetPackageName() == other.GetPackageName()
}

func (cls *Class) IsSubClassOf(other *Class) bool {
	this := cls
	for this != nil {
		if this.superClass == other {
			return true
		}
		this = this.superClass
	}
	return false
}

func (cls *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(cls)
}

func (cls *Class) AccessFlags() uint16 {
	return cls.accessFlags
}

func (cls *Class) NewObject() *Object {
	return newObject(cls)
}

func (cls *Class) RuntimeConstantPool() *RuntimeConstantPool {
	return cls.rtcp
}

func (cls *Class) StaticVars() Slots {
	return cls.staticVars
}

// other cast to this
func (cls *Class) isAssignableFrom(other *Class) bool {
	sub, this := other, cls
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

func (cls *Class) isJvObject() bool {
	return cls.name == "java/lang/Object"
}

func (cls *Class) isJvCloneable() bool {
	return cls.name == "java/lang/Cloneable"
}

func (cls *Class) isJvSerializable() bool {
	return cls.name == "java/io/Serializable"
}

func (cls *Class) IsImplements(other *Class) bool {
	this := cls
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

func (cls *Class) isSubInterfaceOf(other *Class) bool {
	for _, intface := range cls.interfaces {
		if intface == other || intface.isSubInterfaceOf(other) {
			return true
		}
	}
	return false
}

func (cls *Class) isSuperInterfaceOf(other *Class) bool {
	return other.isSubInterfaceOf(cls)
}

func (cls *Class) GetMainMethod() *Method {
	return cls.getMethod("main", "([Ljava/lang/String;)V", true)
}

func (cls *Class) getStaticMethod(name, descriptor string) *Method {
	return cls.getMethod(name, descriptor, true)
}

func (cls *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for this := cls; this != nil; this = this.superClass {
		for _, method := range this.methods {
			if method.IsStatic() == isStatic && method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func (cls *Class) SuperClass() *Class {
	return cls.superClass
}

func (cls *Class) Name() string {
	return cls.name
}

func (cls *Class) ClassLoader() *ClassLoader {
	return cls.classLoader
}

func (cls *Class) GetClinitMethod() *Method {
	return cls.getStaticMethod("<clinit>", "()V")
}

func (cls *Class) IsPublic() bool {
	return cls.accessFlags&ACC_PUBLIC > 0
}

func (cls *Class) IsFinal() bool {
	return cls.accessFlags&ACC_FINAL > 0
}

func (cls *Class) IsSuper() bool {
	return cls.accessFlags&ACC_SUPER > 0
}

func (cls *Class) IsInterface() bool {
	return cls.accessFlags&ACC_INTERFACE > 0
}

func (cls *Class) IsAbstract() bool {
	return cls.accessFlags&ACC_ABSTRACT > 0
}

func (cls *Class) IsSynthetic() bool {
	return cls.accessFlags&ACC_SYNTHETIC > 0
}

func (cls *Class) IsAnnotation() bool {
	return cls.accessFlags&ACC_ANNOTATION > 0
}

func (cls *Class) IsEnum() bool {
	return cls.accessFlags&ACC_ENUM > 0
}

func (cls *Class) IsProtected() bool {
	return cls.accessFlags&ACC_PROTECTED > 0
}

func (cls *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[cls.name]
	return ok
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

func (cls *Class) IsArray() bool {
	return cls.name[0] == '['
}

// getElemName, for example, int[] -> [I, String[] -> [Ljava/lang/String;
func (cls *Class) ArrayClass() *Class {
	arrClassName := getArrayClassName(cls.name)
	return cls.classLoader.LoadClass(arrClassName)
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

func (cls *Class) ComponentClass() *Class {
	componetClassName := getComponentClassName(cls.name)
	return cls.classLoader.LoadClass(componetClassName)
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

func (cls *Class) getField(name, descriptor string, isStatic bool) *Field {
	for this := cls; this != nil; this = this.superClass {
		for _, field := range this.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (cls *Class) GetInstanceMethod(name, descriptor string) *Method {
	return cls.getMethod(name, descriptor, false)
}

func (cls *Class) GetStaticMethod(name, descriptor string) *Method {
	return cls.getMethod(name, descriptor, true)
}

func (cls *Class) GetStaticRefVar(name, descriptor string) *Object {
	field := cls.getField(name, descriptor, true)
	return cls.staticVars.GetRef(field.slotId)
}

func (cls *Class) SetStaticRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := cls.getField(fieldName, fieldDescriptor, true)
	cls.staticVars.SetRef(field.slotId, ref)
}

func (cls *Class) SourceFile() string {
	return cls.sourceFile
}
