package misc

import (
	"go-jvm/src/ch11/instruction/base"
	"go-jvm/src/ch11/native"
	"go-jvm/src/ch11/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize();
func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().ClassLoader()
	jvSysClass := classLoader.LoadClass("java/lang/System")
	initSysClassMethod := jvSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClassMethod)
}
