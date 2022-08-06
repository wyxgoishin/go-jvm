package lang

import (
	"go-jvm/src/ch10/native"
	"go-jvm/src/ch10/rtda"
	"go-jvm/src/ch10/rtda/heap"
)

func init() {
	native.Register("java/lang/System", "arraycopy",
		"(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

// public static native void arraycopy(Object src, int srcPos, Object dst, int dstPos, int length)
func arraycopy(frame *rtda.Frame) {
	localVars := frame.LocalVars
	src := localVars.GetRef(0)
	srcPos := localVars.GetInt(1)
	dst := localVars.GetRef(2)
	dstPos := localVars.GetInt(3)
	length := localVars.GetInt(4)

	if src == nil || dst == nil {
		panic("java.lang.NullPointerException")
	}
	if !checkArrayCopy(src, dst) {
		panic("java.lang.ArrayStoreException")
	}
	if srcPos < 0 || dstPos < 0 || length < 0 || srcPos+length > src.ArrayLength() || dstPos+length > dst.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}
	heap.ArrayCopy(src, srcPos, dst, dstPos, length)
}

func checkArrayCopy(src, dst *heap.Object) bool {
	srcClass := src.Class()
	dstClass := dst.Class()
	if !srcClass.IsArray() || !dstClass.IsArray() {
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() || dstClass.ComponentClass().IsPrimitive() {
		return srcClass == dstClass
	}
	return srcClass == dstClass
}
