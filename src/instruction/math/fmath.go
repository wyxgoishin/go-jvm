package math

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"math"
)

type FADD struct {
	base.NoOperandsInstruction
}

func (inst *FADD) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopFloat()
	val1 := frame.OperandStack().PopFloat()
	frame.OperandStack().PushFloat(val1 + val2)
}

type FSUB struct {
	base.NoOperandsInstruction
}

func (inst *FSUB) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopFloat()
	val1 := frame.OperandStack().PopFloat()
	frame.OperandStack().PushFloat(val1 - val2)
}

type FMUL struct {
	base.NoOperandsInstruction
}

func (inst *FMUL) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopFloat()
	val1 := frame.OperandStack().PopFloat()
	frame.OperandStack().PushFloat(val1 * val2)
}

type FDIV struct {
	base.NoOperandsInstruction
}

func (inst *FDIV) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopFloat()
	val1 := frame.OperandStack().PopFloat()
	frame.OperandStack().PushFloat(val1 / val2)
}

type FNEG struct {
	base.NoOperandsInstruction
}

func (inst *FNEG) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopFloat()
	frame.OperandStack().PushFloat(-val)
}

type FREM struct {
	base.NoOperandsInstruction
}

func (inst *FREM) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopFloat()
	val1 := frame.OperandStack().PopFloat()
	frame.OperandStack().PushFloat(float32(math.Mod(float64(val1), float64(val2))))
}
