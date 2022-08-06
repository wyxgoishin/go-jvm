package conversion

import (
	"go-jvm/src/ch11/instruction/base"
	"go-jvm/src/ch11/rtda"
	"go-jvm/src/ch11/utils"
	"math"
)

type D2F struct {
	base.NoOperandsInstruction
}

func (d *D2F) Execute(frame *rtda.Frame) {
	dval := frame.OperandStack.PopDouble()
	var fval float32
	if math.IsNaN(dval) {
		fval = 0
	} else if dval > math.MaxFloat32 {
		fval = utils.FLOAT32_INF
	} else if dval < math.SmallestNonzeroFloat32 {
		if dval >= 0 {
			fval = 0
		} else {
			fval = -0
		}
	} else {
		fval = float32(dval)
	}
	frame.OperandStack.PushFloat(fval)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (d *D2I) Execute(frame *rtda.Frame) {
	dval := frame.OperandStack.PopDouble()
	var ival int32
	if math.IsNaN(dval) {
		ival = 0
	} else if dval < math.MinInt32 {
		ival = math.MinInt32
	} else if dval > math.MaxInt32 {
		ival = math.MaxInt32
	} else {
		ival = int32(dval)
	}
	frame.OperandStack.PushInt(ival)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (d *D2L) Execute(frame *rtda.Frame) {
	dval := frame.OperandStack.PopDouble()
	var lval int64
	if math.IsNaN(dval) {
		lval = 0
	} else if dval < math.MinInt64 {
		lval = math.MinInt64
	} else if dval > math.MaxInt64 {
		lval = math.MaxInt64
	} else {
		lval = int64(dval)
	}
	frame.OperandStack.PushLong(lval)
}
