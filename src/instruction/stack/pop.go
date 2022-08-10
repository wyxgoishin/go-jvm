package stack

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

func (inst *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (inst *POP2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}
