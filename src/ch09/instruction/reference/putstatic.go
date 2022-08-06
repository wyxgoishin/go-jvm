package reference

import (
	"go-jvm/src/ch09/instruction/base"
	"go-jvm/src/ch09/rtda"
	"go-jvm/src/ch09/rtda/heap"
)

type PUT_STATIC struct {
	base.Index16Struction
}

func (p *PUT_STATIC) Execute(frame *rtda.Frame) {
	curMethod := frame.Method()
	curClass := curMethod.Class()
	rtcp := curClass.RuntimeConstantPool()
	fieldRef := rtcp.GetConstant(p.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if curClass != class || curMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		// ToDo
	}
}
