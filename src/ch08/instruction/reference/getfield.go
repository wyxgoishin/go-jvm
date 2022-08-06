package reference

import (
	"go-jvm/src/ch08/instruction/base"
	"go-jvm/src/ch08/rtda"
	"go-jvm/src/ch08/rtda/heap"
)

type GET_FIELD struct {
	base.Index16Struction
}

func (g *GET_FIELD) Execute(frame *rtda.Frame) {
	rtcp := frame.Method().Class().RuntimeConstantPool()
	fieldRef := rtcp.GetConstant(g.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	stack := frame.OperandStack
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'L':
		stack.PushRef(slots.GetRef(slotId))
	}
}
