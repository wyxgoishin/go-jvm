package heap

import (
	"go-jvm/src/classfile"
)

type SymbolicRef struct {
	rtcp      *RuntimeConstantPool
	className string
	class     *Class
}

func (ref *SymbolicRef) ResolvedClass() *Class {
	if ref.class == nil {
		ref.resolveClassRef()
	}
	return ref.class
}

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-5.html#jvms-5.4.3.5
func (ref *SymbolicRef) resolveClassRef() {
	d := ref.rtcp.class
	c := d.classLoader.LoadClass(ref.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	ref.class = c
}

type ClassRef struct {
	SymbolicRef
}

func newClassRef(rtcp *RuntimeConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	classRef := new(ClassRef)
	classRef.rtcp = rtcp
	classRef.className = classInfo.Name()
	return classRef
}

type MemberRef struct {
	SymbolicRef
	name       string
	descriptor string
}

func (ref *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberInfo) {
	ref.className = refInfo.ClassName()
	ref.name, ref.descriptor = refInfo.NameAndDescriptor()
}

func (ref *MemberRef) Name() string {
	return ref.name
}

func (ref *MemberRef) Descriptor() string {
	return ref.descriptor
}

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(rtcp *RuntimeConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	fieldRef := new(FieldRef)
	fieldRef.rtcp = rtcp
	fieldRef.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return fieldRef
}

func (ref *FieldRef) ResolvedField() *Field {
	if ref.field == nil {
		ref.resolveField()
	}
	return ref.field
}

func (ref *FieldRef) resolveField() {
	d := ref.rtcp.class
	c := ref.ResolvedClass()
	field := lookupField(c, ref.name, ref.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	ref.field = field
}

func lookupField(class *Class, name, descriptor string) *Field {
	for _, field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, intface := range class.interfaces {
		if field := lookupField(intface, name, descriptor); field != nil {
			return field
		}
	}
	if class.superClass != nil {
		return lookupField(class.superClass, name, descriptor)
	}
	return nil
}

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(rtcp *RuntimeConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	methodRef := new(MethodRef)
	methodRef.rtcp = rtcp
	methodRef.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return methodRef
}

func (ref *MethodRef) ResolvedMethod() *Method {
	if ref.method == nil {
		ref.resolveMethod()
	}
	return ref.method
}

func (ref *MethodRef) resolveMethod() {
	dstClass := ref.rtcp.class
	class := ref.ResolvedClass()
	if class.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(class, ref.name, ref.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(dstClass) {
		panic("java.lang.IllegalAccessError")
	}
	ref.method = method
}

func lookupMethod(class *Class, name string, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}

func lookupMethodInInterfaces(interfaces []*Class, name string, descriptor string) *Method {
	for _, intface := range interfaces {
		for _, method := range intface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := lookupMethodInInterfaces(intface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}

func LookupMethodInClass(class *Class, name string, descriptor string) *Method {
	for this := class; this != nil; this = this.superClass {
		for _, method := range this.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(rtcp *RuntimeConstantPool, refInfo *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	interfaceMethodRef := new(InterfaceMethodRef)
	interfaceMethodRef.rtcp = rtcp
	interfaceMethodRef.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return interfaceMethodRef
}

func (ref *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if ref.method == nil {
		ref.resolveMethod()
	}
	return ref.method
}

func (ref *InterfaceMethodRef) resolveMethod() {
	d := ref.rtcp.class
	c := ref.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupInterfaceMethod(c, ref.name, ref.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	ref.method = method
}

func lookupInterfaceMethod(c *Class, name string, descriptor string) *Method {
	for _, method := range c.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(c.interfaces, name, descriptor)
}
