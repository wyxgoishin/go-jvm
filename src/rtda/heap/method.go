package heap

import (
	"go-jvm/src/classfile"
	. "go-jvm/src/utils"
)

type Method struct {
	ClassMember
	maxStack        uint16
	maxLocals       uint16
	code            []byte
	argSlotCount    uint
	exceptionTable  ExceptionTable
	lineNumberTable *classfile.LineNumberTableAttribute
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
		method.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(), method.class.rtcp)
		method.lineNumberTable = memberInfo.LineNumberTableAttribute()
	}
}

func (method *Method) FindExceptionHandler(class *Class, pc int) int {
	handler := method.exceptionTable.findExceptionHandler(class, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}

// 不是所有方法都能查到
func (method *Method) GetLineNumber(pc int) int {
	if method.IsNative() {
		return -2
	}
	if method.lineNumberTable == nil {
		return -1
	}
	return method.lineNumberTable.GetLineNumber(uint(pc))
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

//ToDo: ShimMethod

var (
	_shimClass  = &Class{name: "~shim"}
	_returnCode = []byte{0xb1} // return
	_athrowCode = []byte{0xbf} // athrow

	_returnMethod = &Method{
		ClassMember: ClassMember{
			accessFlags: ACC_STATIC,
			name:        "<return>",
			class:       _shimClass,
		},
		code: _returnCode,
	}

	_athrowMethod = &Method{
		ClassMember: ClassMember{
			accessFlags: ACC_STATIC,
			name:        "<athrow>",
			class:       _shimClass,
		},
		code: _athrowCode,
	}
)

func ShimReturnMethod() *Method {
	return _returnMethod
}
