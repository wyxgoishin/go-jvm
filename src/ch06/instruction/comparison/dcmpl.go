package comparison

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
)

type DCMPL struct {
	base.NoOperandsInstruction
}

func (d *DCMPL) Execute(frame *rtda.Frame) {
	dcmpl(frame, false)
}

type DCMPG struct {
	base.NoOperandsInstruction
}

func (d *DCMPG) Execute(frame *rtda.Frame) {
	dcmpl(frame, true)
}

func dcmpl(frame *rtda.Frame, gFlag bool) {
	val2 := frame.OperandStack.PopDouble()
	val1 := frame.OperandStack.PopDouble()
	if val1 > val2 || gFlag {
		frame.OperandStack.PushInt(1)
	} else if val1 < val2 || !gFlag {
		frame.OperandStack.PushInt(-1)
	} else {
		frame.OperandStack.PushInt(0)
	}
}
