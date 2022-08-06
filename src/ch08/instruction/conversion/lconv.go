package conversion

import (
	"go-jvm/src/ch08/instruction/base"
	"go-jvm/src/ch08/rtda"
)

type L2D struct {
	base.NoOperandsInstruction
}

func (l *L2D) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopLong()
	frame.OperandStack.PushDouble(float64(val))
}

type L2F struct {
	base.NoOperandsInstruction
}

func (l *L2F) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopLong()
	frame.OperandStack.PushFloat(float32(val))
}

type L2I struct {
	base.NoOperandsInstruction
}

func (l *L2I) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopLong()
	frame.OperandStack.PushInt(int32(val))
}
