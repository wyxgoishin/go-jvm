package conversion

import (
	"go-jvm/src/ch08/instruction/base"
	"go-jvm/src/ch08/rtda"
	"go-jvm/src/ch08/utils"
	"math"
)

type I2B struct {
	base.NoOperandsInstruction
}

func (i *I2B) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val > math.MaxInt8 {
		val = math.MaxInt8
	} else if val < math.MinInt8 {
		val = math.MinInt8
	}
	frame.OperandStack.PushInt(val)
}

type I2C struct {
	base.NoOperandsInstruction
}

func (c *I2C) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val > utils.MAX_CHAR {
		val = utils.MAX_CHAR
	} else if val < 0 {
		val = 0
	}
	frame.OperandStack.PushInt(val)
}

type I2D struct {
	base.NoOperandsInstruction
}

func (i *I2D) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushDouble(float64(frame.OperandStack.PopInt()))
}

type I2F struct {
	base.NoOperandsInstruction
}

func (i *I2F) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushFloat(float32(frame.OperandStack.PopInt()))
}

type I2L struct {
	base.NoOperandsInstruction
}

func (i *I2L) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushLong(int64(frame.OperandStack.PopInt()))
}

type I2S struct {
	base.NoOperandsInstruction
}

func (i *I2S) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val > math.MaxInt16 {
		val = math.MaxInt16
	} else if val < math.MinInt16 {
		val = math.MinInt16
	}
	frame.OperandStack.PushInt(val)
}
