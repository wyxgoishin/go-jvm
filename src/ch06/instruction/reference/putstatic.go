package reference

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
	"go-jvm/src/ch06/rtda/heap"
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

	if !heap.IsStatic(field.AccessFlags()) {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if heap.IsFinal(field.AccessFlags()) {
		if curClass != class || curMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
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
	case 'L':
		slots.SetRef(slotId, stack.PopRef())
	}
}
