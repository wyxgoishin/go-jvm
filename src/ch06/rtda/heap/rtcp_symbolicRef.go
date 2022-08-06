package heap

import "go-jvm/src/ch06/classfile"

type SymbolicRef struct {
	rtcp      *RuntimeConstantPool
	className string
	class     *Class
}

func (s *SymbolicRef) ResolvedClass() *Class {
	if s.class == nil {
		s.resolveClassRef()
	}
	return s.class
}

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-5.html#jvms-5.4.3.5
func (s *SymbolicRef) resolveClassRef() {
	d := s.rtcp.class
	c := d.loader.LoadClass(s.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	s.class = c
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

func (m *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberInfo) {
	m.className = refInfo.ClassName()
	m.name, m.descriptor = refInfo.NameAndDescriptor()
}

func (m *MemberRef) Name() string {
	return m.name
}

func (m *MemberRef) Descriptor() string {
	return m.descriptor
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

func (f *FieldRef) ResolvedField() *Field {
	if f.field == nil {
		f.resolveField()
	}
	return f.field
}

func (f *FieldRef) resolveField() {
	d := f.rtcp.class
	c := f.ResolvedClass()
	field := lookupField(c, f.name, f.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	f.field = field
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
	methods *Method
}

func newMethodRef(rtcp *RuntimeConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	methodRef := new(MethodRef)
	methodRef.rtcp = rtcp
	methodRef.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return methodRef
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
