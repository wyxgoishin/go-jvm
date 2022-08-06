package reference

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
	"go-jvm/src/ch06/rtda/heap"
)

type NEW struct {
	base.Index16Struction
}

func (n *NEW) Execute(frame *rtda.Frame) {
	rtcp := frame.Method().Class().RuntimeConstantPool()
	classRef := rtcp.GetConstant(n.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if heap.IsInterface(class.AccessFlags()) || heap.IsAbstract(class.AccessFlags()) {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack.PushRef(ref)
}
