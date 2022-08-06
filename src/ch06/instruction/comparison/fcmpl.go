package comparison

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
	"go-jvm/src/ch06/utils"
)

type FCMPG struct {
	base.NoOperandsInstruction
}

func (f *FCMPG) Execute(frame *rtda.Frame) {
	fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (f *FCMPL) Execute(frame *rtda.Frame) {
	fcmp(frame, false)
}

func fcmp(frame *rtda.Frame, gFlag bool) {
	val2 := frame.OperandStack.PopFloat()
	val1 := frame.OperandStack.PopFloat()
	if utils.IsNaN(val1) || utils.IsNaN(val2) {
		if gFlag {
			frame.OperandStack.PushFloat(1)
		} else {
			frame.OperandStack.PushFloat(-1)
		}
	} else if val1 == val2 {
		frame.OperandStack.PushFloat(0)
	} else if val1 > val2 {
		frame.OperandStack.PushFloat(1)
	} else {
		frame.OperandStack.PushFloat(-1)
	}
}
