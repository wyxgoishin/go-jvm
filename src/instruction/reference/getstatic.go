package reference

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"go-jvm/src/rtda/heap"
)

type GET_STATIC struct {
	base.Index16Struction
}

func (inst *GET_STATIC) Execute(frame *rtda.Frame) {
	rtcp := frame.Method().Class().RuntimeConstantPool()
	fieldRef := rtcp.GetConstant(inst.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	staticVars := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(staticVars.GetInt(slotId))
	case 'F':
		stack.PushFloat(staticVars.GetFloat(slotId))
	case 'D':
		stack.PushDouble(staticVars.GetDouble(slotId))
	case 'J':
		stack.PushLong(staticVars.GetLong(slotId))
	case 'L', '[':
		stack.PushRef(staticVars.GetRef(slotId))
	}
}
