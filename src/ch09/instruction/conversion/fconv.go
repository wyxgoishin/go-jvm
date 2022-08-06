package conversion

import (
	"go-jvm/src/ch09/instruction/base"
	"go-jvm/src/ch09/rtda"
	"math"
)

type F2D struct {
	base.NoOperandsInstruction
}

func (f *F2D) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushDouble(float64(frame.OperandStack.PopFloat()))
}

type F2I struct {
	base.NoOperandsInstruction
}

func (f *F2I) Execute(frame *rtda.Frame) {
	fval := frame.OperandStack.PopFloat()
	var ival int32
	if math.IsNaN(float64(fval)) {
		ival = 0
	} else if fval > math.MaxInt32 {
		ival = math.MaxInt32
	} else if fval < math.MinInt32 {
		ival = math.MinInt32
	} else {
		ival = int32(fval)
	}
	frame.OperandStack.PushInt(ival)
}

type F2L struct {
	base.NoOperandsInstruction
}

func (f *F2L) Execute(frame *rtda.Frame) {
	fval := frame.OperandStack.PopFloat()
	var lval int64
	if math.IsNaN(float64(fval)) {
		lval = 0
	} else if fval > math.MaxInt64 {
		lval = math.MaxInt64
	} else if fval < math.MinInt64 {
		lval = math.MinInt64
	} else {
		lval = int64(fval)
	}
	frame.OperandStack.PushLong(lval)
}
