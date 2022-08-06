package store

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

type AASTORE struct {
	base.NoOperandsInstruction
}

func (a *AASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopRef()
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	arrRef.Refs()[index] = val
}

type BASTORE struct {
	base.NoOperandsInstruction
}

func (b *BASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	arrRef.Bytes()[index] = int8(val)
}

type CASTORE struct {
	base.NoOperandsInstruction
}

func (c *CASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	arrRef.Chars()[index] = uint16(val)
}

type DASTORE struct {
	base.NoOperandsInstruction
}

func (d *DASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopDouble()
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	arrRef.Doubles()[index] = val
}

type FASTORE struct {
	base.NoOperandsInstruction
}

func (f *FASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopFloat()
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	arrRef.Floats()[index] = val
}

type IASTORE struct {
	base.NoOperandsInstruction
}

func (i *IASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	arrRef.Ints()[index] = val
}

type LASTORE struct {
	base.NoOperandsInstruction
}

func (l *LASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopLong()
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	arrRef.Longs()[index] = val
}

type SASTORE struct {
	base.NoOperandsInstruction
}

func (s *SASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	index := frame.OperandStack.PopInt()
	arrRef := frame.OperandStack.PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(uint(arrLen), uint(index))
	arrRef.Shorts()[index] = int16(val)
}
