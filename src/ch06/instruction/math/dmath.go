package math

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
	"math"
)

type DADD struct {
	base.NoOperandsInstruction
}

func (d *DADD) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopDouble()
	val2 := frame.OperandStack.PopDouble()
	frame.OperandStack.PushDouble(val1 + val2)
}

type DSUB struct {
	base.NoOperandsInstruction
}

func (d *DSUB) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopDouble()
	val2 := frame.OperandStack.PopDouble()
	frame.OperandStack.PushDouble(val1 - val2)
}

type DMUL struct {
	base.NoOperandsInstruction
}

func (d *DMUL) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopDouble()
	val2 := frame.OperandStack.PopDouble()
	frame.OperandStack.PushDouble(val1 * val2)
}

type DDIV struct {
	base.NoOperandsInstruction
}

func (d *DDIV) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopDouble()
	val2 := frame.OperandStack.PopDouble()
	frame.OperandStack.PushDouble(val1 / val2)
}

type DNEG struct {
	base.NoOperandsInstruction
}

func (d *DNEG) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushDouble(-frame.OperandStack.PopDouble())
}

type DREM struct {
	base.NoOperandsInstruction
}

func (d *DREM) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopDouble()
	val2 := frame.OperandStack.PopDouble()
	frame.OperandStack.PushDouble(math.Mod(val1, val2))
}
