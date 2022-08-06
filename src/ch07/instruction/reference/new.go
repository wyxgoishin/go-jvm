package reference

import (
	"go-jvm/src/ch07/instruction/base"
	"go-jvm/src/ch07/rtda"
	"go-jvm/src/ch07/rtda/heap"
)

type NEW struct {
	base.Index16Struction
}

func (n *NEW) Execute(frame *rtda.Frame) {
	rtcp := frame.Method().Class().RuntimeConstantPool()
	classRef := rtcp.GetConstant(n.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	ref := class.NewObject()
	frame.OperandStack.PushRef(ref)
}
