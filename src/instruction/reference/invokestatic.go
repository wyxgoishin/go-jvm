package reference

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"go-jvm/src/rtda/heap"
	"strings"
)

type INVOKE_STATIC struct {
	base.Index16Struction
}

func (inst *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	rtcp := frame.Method().Class().RuntimeConstantPool()
	methodRef := rtcp.GetConstant(inst.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if strings.Contains(resolvedMethod.Name(), "toUnsignedString0") {
		println("debug")
	}
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	class := resolvedMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	base.InvokeMethod(frame, resolvedMethod)
}
