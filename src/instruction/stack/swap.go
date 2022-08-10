package stack

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (inst *SWAP) Execute(frame *rtda.Frame) {
	slot1 := frame.OperandStack().PopSlot()
	slot2 := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot2)
	frame.OperandStack().PushSlot(slot1)
}
