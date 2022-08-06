package reference

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
)

type INVOKE_SPECIAL struct {
	base.Index16Struction
}

//ToDo: currently only hack
func (i *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack.PopRef()
}
