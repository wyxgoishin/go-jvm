package reference

import (
	"go-jvm/src/ch08/instruction/base"
	"go-jvm/src/ch08/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (a ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	arrRef := frame.OperandStack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.ArrayLength()
	frame.OperandStack.PushInt(arrLen)
}
