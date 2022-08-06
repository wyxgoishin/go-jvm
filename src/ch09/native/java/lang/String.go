package lang

import (
	"go-jvm/src/ch09/native"
	"go-jvm/src/ch09/rtda"
	"go-jvm/src/ch09/rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	interned := heap.InternedString(this)
	frame.OperandStack.PushRef(interned)
}
