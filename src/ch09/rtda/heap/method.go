package heap

import (
	"go-jvm/src/ch09/classfile"
	. "go-jvm/src/ch09/utils"
)

type Method struct {
	ClassMember
	maxStack     uint16
	maxLocals    uint16
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := new(Method)
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	methodDescriptor := parseMethodDescriptor(cfMethod.Descriptor())
	method.calcArgSlotCount(methodDescriptor.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(methodDescriptor.returnType)
	}
	return method
}

func (method *Method) injectCodeAttribute(returnType string) {
	// ToDo:
	method.maxStack = 4
	method.maxLocals = uint16(method.argSlotCount)
	switch returnType[0] {
	case 'V':
		method.code = []byte{String2OpCode["impdep1"], String2OpCode["_return"]}
	case 'D':
		method.code = []byte{String2OpCode["impdep1"], String2OpCode["dreturn"]}
	case 'F':
		method.code = []byte{String2OpCode["impdep1"], String2OpCode["freturn"]}
	case 'J':
		method.code = []byte{String2OpCode["impdep1"], String2OpCode["lreturn"]}
	case 'L':
		method.code = []byte{String2OpCode["impdep1"], String2OpCode["areturn"]}
	default:
		method.code = []byte{String2OpCode["impdep1"], String2OpCode["ireturn"]}
	}
}

func (method *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		method.argSlotCount++
		if paramType == "J" || paramType == "D" {
			method.argSlotCount++
		}
	}
	if !method.IsStatic() {
		method.argSlotCount++
	}
}

func (method *Method) copyAttributes(memberInfo *classfile.MemberInfo) {
	if codeAttr := memberInfo.CodeAttribute(); codeAttr != nil {
		method.maxLocals = codeAttr.MaxLocals()
		method.maxStack = codeAttr.MaxStack()
		method.code = codeAttr.Code()
	}
}

func (method *Method) Class() *Class {
	return method.class
}

func (method *Method) MaxStack() uint {
	return uint(method.maxStack)
}

func (method *Method) MaxLocals() uint {
	return uint(method.maxLocals)
}

func (method *Method) Code() []byte {
	return method.code
}

func (method *Method) ArgSlotCount() uint {
	return method.argSlotCount
}

func (method *Method) IsSynchronized() bool {
	return method.accessFlags&ACC_SYNCHRONIZED > 0
}

func (method *Method) IsBridge() bool {
	return method.accessFlags&ACC_BRIDGE > 0
}

func (method *Method) IsVarArgs() bool {
	return method.accessFlags&ACC_VARARGS > 0
}

func (method *Method) IsNative() bool {
	return method.accessFlags&ACC_NATIVE > 0
}

func (method *Method) IsAbstract() bool {
	return method.accessFlags&ACC_ABSTRACT > 0
}

func (method *Method) IsStrict() bool {
	return method.accessFlags&ACC_STRICT > 0
}
