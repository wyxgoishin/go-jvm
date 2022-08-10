package reference

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"go-jvm/src/rtda/heap"
	"go-jvm/src/utils"
)

type NEW struct {
	base.Index16Struction
}

func (inst *NEW) Execute(frame *rtda.Frame) {
	rtcp := frame.Method().Class().RuntimeConstantPool()
	classRef := rtcp.GetConstant(inst.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}

const (
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

type NEW_ARRAY struct {
	atype uint8
}

func (inst *NEW_ARRAY) FetchOperands(reader *utils.ByteCodeReader) {
	inst.atype = reader.ReadUint8()
}

func (inst *NEW_ARRAY) Execute(frame *rtda.Frame) {
	count := frame.OperandStack().PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	classLoader := frame.Method().Class().ClassLoader()
	arrClass := getPrimitiveArrayClass(classLoader, inst.atype)
	arr := arrClass.NewArray(uint(count))
	frame.OperandStack().PushRef(arr)
}

func getPrimitiveArrayClass(classLoader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return classLoader.LoadClass("[Z")
	case AT_BYTE:
		return classLoader.LoadClass("[B")
	case AT_CHAR:
		return classLoader.LoadClass("[C")
	case AT_SHORT:
		return classLoader.LoadClass("[S")
	case AT_INT:
		return classLoader.LoadClass("[I")
	case AT_FLOAT:
		return classLoader.LoadClass("[F")
	case AT_LONG:
		return classLoader.LoadClass("[J")
	case AT_DOUBLE:
		return classLoader.LoadClass("[D")
	default:
		panic("Unexpected array elem type: " + string(atype))
	}
}

type ANEW_ARRAY struct {
	base.Index16Struction
}

func (inst *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	rtcp := frame.Method().Class().RuntimeConstantPool()
	classRef := rtcp.GetConstant(inst.Index).(*heap.ClassRef)
	elemClass := classRef.ResolvedClass()
	count := frame.OperandStack().PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrClass := elemClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	frame.OperandStack().PushRef(arr)
}

type MULTI_ANEW_ARRAY struct {
	index     uint16
	dimension uint8
}

func (inst *MULTI_ANEW_ARRAY) FetchOperands(reader *utils.ByteCodeReader) {
	inst.index = reader.ReadUint16()
	inst.dimension = reader.ReadUint8()
}

func (inst *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	rtcp := frame.Method().Class().RuntimeConstantPool()
	classRef := rtcp.GetConstant(uint(inst.index)).(*heap.ClassRef)
	elemClass := classRef.ResolvedClass()
	counts := popAndCheckCount(frame.OperandStack(), int8(inst.dimension))
	arr := newMultiDimensionArray(counts, elemClass)
	frame.OperandStack().PushRef(arr)
}

func popAndCheckCount(stack *rtda.OperandStack, dimension int8) []int32 {
	counts := make([]int32, dimension)
	for i := dimension - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}

func newMultiDimensionArray(counts []int32, class *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := class.NewArray(count)
	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionArray(counts[i:], class.ComponentClass())
		}
	}
	return arr
}
