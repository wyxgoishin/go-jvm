package load

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"go-jvm/src/rtda/heap"
)

func checkNotNull(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndexInBound(len int, index int32) {
	if index < 0 || int(index) >= len {
		panic("java.lang.ArrayIndexOutOfBoundException")
	}
}

type AALOAD struct {
	base.NoOperandsInstruction
}

func (inst *AALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	arrRef := frame.OperandStack().PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(arrLen, index)
	aval := arrRef.Refs()[index]
	frame.OperandStack().PushRef(aval)
}

type BALOAD struct {
	base.NoOperandsInstruction
}

func (inst *BALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	arrRef := frame.OperandStack().PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(arrLen, index)
	bval := int32(arrRef.Bytes()[index])
	frame.OperandStack().PushInt(bval)
}

type CALOAD struct {
	base.NoOperandsInstruction
}

func (inst *CALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	arrRef := frame.OperandStack().PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(arrLen, index)
	cval := int32(arrRef.Chars()[index])
	frame.OperandStack().PushInt(cval)
}

type DALOAD struct {
	base.NoOperandsInstruction
}

func (inst *DALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	arrRef := frame.OperandStack().PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(arrLen, index)
	dval := arrRef.Doubles()[index]
	frame.OperandStack().PushDouble(dval)
}

type FALOAD struct {
	base.NoOperandsInstruction
}

func (inst *FALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	arrRef := frame.OperandStack().PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(arrLen, index)
	fval := arrRef.Floats()[index]
	frame.OperandStack().PushFloat(fval)
}

type IALOAD struct {
	base.NoOperandsInstruction
}

func (i *IALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	arrRef := frame.OperandStack().PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(arrLen, index)
	ival := arrRef.Ints()[index]
	frame.OperandStack().PushInt(ival)
}

type LALOAD struct {
	base.NoOperandsInstruction
}

func (inst *LALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	arrRef := frame.OperandStack().PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(arrLen, index)
	lval := arrRef.Longs()[index]
	frame.OperandStack().PushLong(lval)
}

type SALOAD struct {
	base.NoOperandsInstruction
}

func (inst *SALOAD) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	arrRef := frame.OperandStack().PopRef()
	checkNotNull(arrRef)
	arrLen := arrRef.ArrayLength()
	checkIndexInBound(arrLen, index)
	sval := int32(arrRef.Shorts()[index])
	frame.OperandStack().PushInt(sval)
}
