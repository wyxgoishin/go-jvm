package stack

import (
	"go-jvm/src/ch10/instruction/base"
	"go-jvm/src/ch10/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

func (p *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack.PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (p *POP2) Execute(frame *rtda.Frame) {
	frame.OperandStack.PopSlot()
	frame.OperandStack.PopSlot()
}
