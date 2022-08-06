package reference

import (
	"go-jvm/src/ch08/instruction/base"
	"go-jvm/src/ch08/rtda"
	"go-jvm/src/ch08/rtda/heap"
)

type INSTANCEOF struct {
	base.Index16Struction
}

func (i *INSTANCEOF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	rtcp := frame.Method().Class().RuntimeConstantPool()
	classRef := rtcp.GetConstant(i.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
