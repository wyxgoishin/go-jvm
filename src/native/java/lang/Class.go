package lang

import (
	"go-jvm/src/native"
	"go-jvm/src/rtda"
	"go-jvm/src/rtda/heap"
)

func init() {
	native.Register("java/lang/Class", "getPrimitiveClass",
		"(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0",
		"()Ljava/lang/String;", getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0",
		"(Ljava/lang/Class;)Z", desiredAssertionStatus)
}

// static native Class<?> getPrimitiveClass(String name)
func getPrimitiveClass(frame *rtda.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)
	classloader := frame.Method().Class().ClassLoader()
	class := classloader.LoadClass(name).JvClass()
	frame.OperandStack().PushRef(class)
}

// private native String getName0()
func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)
	name := class.JvName()
	nameObj := heap.JvString(class.ClassLoader(), name)
	frame.OperandStack().PushRef(nameObj)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
func desiredAssertionStatus(frame *rtda.Frame) {
	//ToDo: Implementing me
	frame.OperandStack().PushBoolean(false)
}
