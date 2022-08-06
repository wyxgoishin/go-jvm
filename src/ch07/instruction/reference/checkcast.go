package reference

import (
	"go-jvm/src/ch07/instruction/base"
	"go-jvm/src/ch07/rtda"
	"go-jvm/src/ch07/rtda/heap"
)

type CHECK_CAST struct {
	base.Index16Struction
}

func (c *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	rtcp := frame.Method().Class().RuntimeConstantPool()
	classRef := rtcp.GetConstant(c.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
