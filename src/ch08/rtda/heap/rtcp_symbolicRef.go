package heap

import "go-jvm/src/ch08/classfile"

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
	c := d.classLoader.LoadClass(s.className)
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
	method *Method
}

func newMethodRef(rtcp *RuntimeConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	methodRef := new(MethodRef)
	methodRef.rtcp = rtcp
	methodRef.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return methodRef
}

func (m *MethodRef) ResolvedMethod() *Method {
	if m.method == nil {
		m.resolveMethod()
	}
	return m.method
}

func (m *MethodRef) resolveMethod() {
	d := m.rtcp.class
	c := m.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, m.name, m.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	m.method = method
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

func (i *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if i.method == nil {
		i.resolveMethod()
	}
	return i.method
}

func (i *InterfaceMethodRef) resolveMethod() {
	d := i.rtcp.class
	c := i.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupInterfaceMethod(c, i.name, i.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	i.method = method
}

func lookupInterfaceMethod(c *Class, name string, descriptor string) *Method {
	for _, method := range c.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(c.interfaces, name, descriptor)
}
