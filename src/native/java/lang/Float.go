package lang

import (
	"go-jvm/src/native"
	"go-jvm/src/rtda"
	"math"
)

func init() {
	native.Register("java/lang/Float", "floatToRawIntBits", "(F)I", floatToRawIntBits)
}

func floatToRawIntBits(frame *rtda.Frame) {
	val := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(val)
	frame.OperandStack().PushInt(int32(bits))
}
