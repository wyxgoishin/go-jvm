package math

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"math"
)

type DADD struct {
	base.NoOperandsInstruction
}

func (inst *DADD) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopDouble()
	val1 := frame.OperandStack().PopDouble()
	frame.OperandStack().PushDouble(val2 + val1)
}

type DSUB struct {
	base.NoOperandsInstruction
}

func (inst *DSUB) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopDouble()
	val1 := frame.OperandStack().PopDouble()
	frame.OperandStack().PushDouble(val1 - val2)
}

type DMUL struct {
	base.NoOperandsInstruction
}

func (inst *DMUL) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopDouble()
	val1 := frame.OperandStack().PopDouble()
	frame.OperandStack().PushDouble(val1 * val2)
}

type DDIV struct {
	base.NoOperandsInstruction
}

func (inst *DDIV) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopDouble()
	val1 := frame.OperandStack().PopDouble()
	frame.OperandStack().PushDouble(val1 / val2)
}

type DNEG struct {
	base.NoOperandsInstruction
}

func (inst *DNEG) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopDouble()
	frame.OperandStack().PushDouble(-val)
}

type DREM struct {
	base.NoOperandsInstruction
}

func (inst *DREM) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopDouble()
	val1 := frame.OperandStack().PopDouble()
	frame.OperandStack().PushDouble(math.Mod(val1, val2))
}
