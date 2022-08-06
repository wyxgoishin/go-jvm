package lang

import (
	"go-jvm/src/ch10/native"
	"go-jvm/src/ch10/rtda"
	"math"
)

func init() {
	native.Register("java/lang/Double", "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register("java/lang/Double", "longBitsToDouble", "(J)D", longBitsToDouble)
}

func doubleToRawLongBits(frame *rtda.Frame) {
	val := frame.LocalVars.GetDouble(0)
	bits := math.Float64bits(val)
	frame.OperandStack.PushLong(int64(bits))
}

func longBitsToDouble(frame *rtda.Frame) {
	lval := frame.LocalVars.GetLong(0)
	// ToDo: Implementing me
	dval := math.Float64frombits(uint64(lval))
	frame.OperandStack.PushDouble(dval)
}
