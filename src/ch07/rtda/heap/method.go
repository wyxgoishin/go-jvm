package heap

import "go-jvm/src/ch07/classfile"

type Method struct {
	ClassMember
	maxStack     uint16
	maxLocals    uint16
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, method := range cfMethods {
		methods[i] = new(Method)
		methods[i].class = class
		methods[i].copyMemberInfo(method)
		methods[i].copyAttributes(method)
		methods[i].calcArgSlotCount()
	}
	return methods
}

func (m *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(m.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		m.argSlotCount++
		if paramType == "J" || paramType == "D" {
			m.argSlotCount++
		}
	}
	if !m.IsStatic() {
		m.argSlotCount++
	}
}

func (m *Method) copyAttributes(memberInfo *classfile.MemberInfo) {
	if codeAttr := memberInfo.CodeAttribute(); codeAttr != nil {
		m.maxLocals = codeAttr.MaxLocals()
		m.maxStack = codeAttr.MaxStack()
		m.code = codeAttr.Code()
	}
}

func (m *Method) Class() *Class {
	return m.class
}

func (m *Method) MaxStack() uint {
	return uint(m.maxStack)
}

func (m *Method) MaxLocals() uint {
	return uint(m.maxLocals)
}

func (m *Method) Code() []byte {
	return m.code
}

func (m *Method) ArgSlotCount() uint {
	return m.argSlotCount
}

func (m *Method) IsSynchronized() bool {
	return m.accessFlags&ACC_SYNCHRONIZED > 0
}

func (m *Method) IsBridge() bool {
	return m.accessFlags&ACC_BRIDGE > 0
}

func (m *Method) IsVarArgs() bool {
	return m.accessFlags&ACC_VARARGS > 0
}

func (m *Method) IsNative() bool {
	return m.accessFlags&ACC_NATIVE > 0
}

func (m *Method) IsAbstract() bool {
	return m.accessFlags&ACC_ABSTRACT > 0
}

func (m *Method) IsStrict() bool {
	return m.accessFlags&ACC_STRICT > 0
}
