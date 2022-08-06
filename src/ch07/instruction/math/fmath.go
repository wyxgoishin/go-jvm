package math

import (
	"go-jvm/src/ch07/instruction/base"
	"go-jvm/src/ch07/rtda"
	"math"
)

type FADD struct {
	base.NoOperandsInstruction
}

func (f *FADD) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopFloat()
	val2 := frame.OperandStack.PopFloat()
	frame.OperandStack.PushFloat(val1 + val2)
}

type FSUB struct {
	base.NoOperandsInstruction
}

func (f *FSUB) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopFloat()
	val2 := frame.OperandStack.PopFloat()
	frame.OperandStack.PushFloat(val2 - val1)
}

type FMUL struct {
	base.NoOperandsInstruction
}

func (f *FMUL) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopFloat()
	val2 := frame.OperandStack.PopFloat()
	frame.OperandStack.PushFloat(val1 * val2)
}

type FDIV struct {
	base.NoOperandsInstruction
}

func (f *FDIV) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopFloat()
	val2 := frame.OperandStack.PopFloat()
	frame.OperandStack.PushFloat(val2 / val1)
}

type FNEG struct {
	base.NoOperandsInstruction
}

func (f *FNEG) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushFloat(-frame.OperandStack.PopFloat())
}

type FREM struct {
	base.NoOperandsInstruction
}

func (f *FREM) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopFloat()
	val2 := frame.OperandStack.PopFloat()
	frame.OperandStack.PushFloat(float32(math.Mod(float64(val2), float64(val1))))
}
