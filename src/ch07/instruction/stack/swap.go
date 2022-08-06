package stack

import (
	"go-jvm/src/ch07/instruction/base"
	"go-jvm/src/ch07/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (s *SWAP) Execute(frame *rtda.Frame) {
	slot1 := frame.OperandStack.PopSlot()
	slot2 := frame.OperandStack.PopSlot()
	frame.OperandStack.PushSlot(slot2)
	frame.OperandStack.PushSlot(slot1)
}
