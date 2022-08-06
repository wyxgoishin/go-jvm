package misc

import (
	"go-jvm/src/ch10/instruction/base"
	"go-jvm/src/ch10/native"
	"go-jvm/src/ch10/rtda"
	"go-jvm/src/ch10/rtda/heap"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize();
func initialize(frame *rtda.Frame) {
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetStaticRefVar("savedProps", "Ljava/util/Properties;")
	key := heap.JvString(vmClass.ClassLoader(), "foo")
	val := heap.JvString(vmClass.ClassLoader(), "bar")
	frame.OperandStack.PushRef(savedProps)
	frame.OperandStack.PushRef(key)
	frame.OperandStack.PushRef(val)
	propsClass := vmClass.ClassLoader().LoadClass("java/util/Properties")
	setPropMethod := propsClass.GetInstanceMethod("setProperty",
		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(frame, setPropMethod)
}
