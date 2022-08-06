package math

import (
	"go-jvm/src/ch08/instruction/base"
	"go-jvm/src/ch08/rtda"
)

type LADD struct {
	base.NoOperandsInstruction
}

func (l *LADD) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopLong()
	val2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val1 + val2)
}

type LSUB struct {
	base.NoOperandsInstruction
}

func (l *LSUB) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopLong()
	val2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val2 - val1)
}

type LMUL struct {
	base.NoOperandsInstruction
}

func (l *LMUL) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopLong()
	val2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val1 * val2)
}

type LDIV struct {
	base.NoOperandsInstruction
}

func (l *LDIV) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopLong()
	val2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val2 / val1)
}

type LNEG struct {
	base.NoOperandsInstruction
}

func (l *LNEG) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushLong(-frame.OperandStack.PopLong())
}

type LREM struct {
	base.NoOperandsInstruction
}

func (l *LREM) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopLong()
	val2 := frame.OperandStack.PopLong()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	frame.OperandStack.PushLong(val2 % val1)
}

type LAND struct {
	base.NoOperandsInstruction
}

func (l *LAND) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopLong()
	val2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val1 & val2)
}

type LOR struct {
	base.NoOperandsInstruction
}

func (l *LOR) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopLong()
	val2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val1 | val2)
}

type LXOR struct {
	base.NoOperandsInstruction
}

func (l *LXOR) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopLong()
	val2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val1 ^ val2)
}

type LSHL struct {
	base.NoOperandsInstruction
}

func (l *LSHL) Execute(frame *rtda.Frame) {
	val1 := uint64(frame.OperandStack.PopLong()) & 63
	val2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val2 << val1)
}

type LSHR struct {
	base.NoOperandsInstruction
}

func (l *LSHR) Execute(frame *rtda.Frame) {
	val1 := uint64(frame.OperandStack.PopLong()) & 63
	val2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val2 >> val1)
}

type LUSHL struct {
	base.NoOperandsInstruction
}

func (l *LUSHL) Execute(frame *rtda.Frame) {
	val1 := uint64(frame.OperandStack.PopLong()) & 63
	val2 := uint64(frame.OperandStack.PopLong())
	frame.OperandStack.PushLong(int64(val2 << val1))
}

type LUSHR struct {
	base.NoOperandsInstruction
}

func (l *LUSHR) Execute(frame *rtda.Frame) {
	val1 := uint64(frame.OperandStack.PopLong()) & 63
	val2 := uint64(frame.OperandStack.PopLong())
	frame.OperandStack.PushLong(int64(val2 >> val1))
}
