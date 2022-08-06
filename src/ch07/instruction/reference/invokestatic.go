package reference

import (
	"go-jvm/src/ch07/instruction/base"
	"go-jvm/src/ch07/rtda"
	"go-jvm/src/ch07/rtda/heap"
)

type INVOKE_STATIC struct {
	base.Index16Struction
}

func (i *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	rtcp := frame.Method().Class().RuntimeConstantPool()
	methodRef := rtcp.GetConstant(i.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
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
