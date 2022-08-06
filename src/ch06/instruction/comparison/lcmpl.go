package comparison

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (l *LCMP) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack.PopLong()
	val1 := frame.OperandStack.PopLong()
	if val1 > val2 {
		frame.OperandStack.PushInt(1)
	} else if val1 == val2 {
		frame.OperandStack.PushInt(0)
	} else {
		frame.OperandStack.PushInt(-1)
	}
}
