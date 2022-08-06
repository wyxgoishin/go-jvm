package load

import (
	"go-jvm/src/ch08/instruction/base"
	"go-jvm/src/ch08/rtda"
	"go-jvm/src/ch08/rtda/heap"
)

func checkNotNull(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndexInBound(len, index uint) {
	if index < 0 || index >= len {
		panic("java.lang.ArrayIndexOutOfBoundException")
	}
}

type AALOAD struct {
	base.NoOperandsInstruction
}

func (a *AALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	frame.OperandStack.PushRef(arrRef.Refs()[index])
}

type BALOAD struct {
	base.NoOperandsInstruction
}

func (b *BALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	frame.OperandStack.PushInt(int32(arrRef.Bytes()[index]))
}

type CALOAD struct {
	base.NoOperandsInstruction
}

func (c *CALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	frame.OperandStack.PushInt(int32(arrRef.Chars()[index]))
}

type DALOAD struct {
	base.NoOperandsInstruction
}

func (d *DALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	frame.OperandStack.PushDouble(float64(arrRef.Doubles()[index]))
}

type FALOAD struct {
	base.NoOperandsInstruction
}

func (f *FALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	frame.OperandStack.PushFloat(float32(arrRef.Floats()[index]))
}

type IALOAD struct {
	base.NoOperandsInstruction
}

func (i *IALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	frame.OperandStack.PushInt(arrRef.Ints()[index])
}

type LALOAD struct {
	base.NoOperandsInstruction
}

func (l *LALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	frame.OperandStack.PushLong(arrRef.Longs()[index])
}

type SALOAD struct {
	base.NoOperandsInstruction
}

func (s *SALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	frame.OperandStack.PushInt(int32(arrRef.Shorts()[index]))
}
