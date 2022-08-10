package constant

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (inst *NOP) Execute(frame *rtda.Frame) {
}
