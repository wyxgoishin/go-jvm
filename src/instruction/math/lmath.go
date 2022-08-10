package math

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type LADD struct {
	base.NoOperandsInstruction
}

func (inst *LADD) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopLong()
	val1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(val1 + val2)
}

type LSUB struct {
	base.NoOperandsInstruction
}

func (inst *LSUB) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopLong()
	val1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(val1 - val2)
}

type LMUL struct {
	base.NoOperandsInstruction
}

func (inst *LMUL) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopLong()
	val1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(val1 * val2)
}

type LDIV struct {
	base.NoOperandsInstruction
}

func (inst *LDIV) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopLong()
	val1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(val1 / val2)
}

type LNEG struct {
	base.NoOperandsInstruction
}

func (inst *LNEG) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(-frame.OperandStack().PopLong())
}

type LREM struct {
	base.NoOperandsInstruction
}

//ToDo: support throw exception
func (inst *LREM) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopLong()
	val1 := frame.OperandStack().PopLong()
	if val1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	frame.OperandStack().PushLong(val1 % val2)
}

type LAND struct {
	base.NoOperandsInstruction
}

func (inst *LAND) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopLong()
	val1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(val1 & val2)
}

type LOR struct {
	base.NoOperandsInstruction
}

func (inst *LOR) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopLong()
	val1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(val1 | val2)
}

type LXOR struct {
	base.NoOperandsInstruction
}

func (inst *LXOR) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopLong()
	val1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(val1 ^ val2)
}

type LSHL struct {
	base.NoOperandsInstruction
}

func (inst *LSHL) Execute(frame *rtda.Frame) {
	val2 := uint64(frame.OperandStack().PopInt()) & 63
	val1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(val1 << val2)
}

type LSHR struct {
	base.NoOperandsInstruction
}

func (inst *LSHR) Execute(frame *rtda.Frame) {
	val2 := uint64(frame.OperandStack().PopInt()) & 63
	val1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(val1 >> val2)
}

type LUSHL struct {
	base.NoOperandsInstruction
}

func (inst *LUSHL) Execute(frame *rtda.Frame) {
	val2 := uint64(frame.OperandStack().PopInt()) & 63
	val1 := uint64(frame.OperandStack().PopLong())
	frame.OperandStack().PushLong(int64(val1 << val2))
}

type LUSHR struct {
	base.NoOperandsInstruction
}

func (inst *LUSHR) Execute(frame *rtda.Frame) {
	val2 := uint64(frame.OperandStack().PopInt()) & 63
	val1 := uint64(frame.OperandStack().PopLong())
	frame.OperandStack().PushLong(int64(val1 >> val2))
}
