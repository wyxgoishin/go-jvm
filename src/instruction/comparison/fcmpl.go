package comparison

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"go-jvm/src/utils"
)

type FCMPG struct {
	base.NoOperandsInstruction
}

func (inst *FCMPG) Execute(frame *rtda.Frame) {
	fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (inst *FCMPL) Execute(frame *rtda.Frame) {
	fcmp(frame, false)
}

func fcmp(frame *rtda.Frame, gFlag bool) {
	val2 := frame.OperandStack().PopFloat()
	val1 := frame.OperandStack().PopFloat()
	if utils.IsNaN(val1) || utils.IsNaN(val2) {
		if gFlag {
			frame.OperandStack().PushInt(1)
		} else {
			frame.OperandStack().PushInt(-1)
		}
	} else if val1 == val2 {
		frame.OperandStack().PushInt(0)
	} else if val1 > val2 {
		frame.OperandStack().PushInt(1)
	} else {
		frame.OperandStack().PushInt(-1)
	}
}
