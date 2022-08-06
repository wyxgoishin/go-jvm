package heap

import "go-jvm/src/ch06/classfile"

type Method struct {
	ClassMember
	maxStack  uint16
	maxLocals uint16
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, method := range cfMethods {
		methods[i] = new(Method)
		methods[i].class = class
		methods[i].copyMemberInfo(method)
		methods[i].copyAttributes(method)
	}
	return methods
}

func (m *Method) copyAttributes(memberInfo *classfile.MemberInfo) {
	if codeAttr := memberInfo.CodeAttribute(); codeAttr != nil {
		m.maxLocals = codeAttr.MaxLocals()
		m.maxStack = codeAttr.MaxLocals()
		m.code = codeAttr.Code()
	}
}

func (m *Method) Class() *Class {
	return m.class
}

func (m *Method) Name() string {
	return m.name
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
