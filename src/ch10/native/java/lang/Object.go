package lang

import (
	"go-jvm/src/ch10/native"
	"go-jvm/src/ch10/rtda"
	"unsafe"
)

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
	native.Register("java/lang/Object", "hashCode", "()I", hashCode)
	native.Register("java/lang/Object", "clone", "()Ljava/lang/Object;", clone)
}

// public final native Class<?> getClass()
func getClass(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	jvClass := this.Class().JvClass()
	frame.OperandStack.PushRef(jvClass)
}

// public int native hashCode();
func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack.PushInt(hash)
}

// protected native Object clone() throws CloneNotSupportedException
func clone(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	cloneable := this.Class().ClassLoader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}
	frame.OperandStack.PushRef(this.Clone())
}
