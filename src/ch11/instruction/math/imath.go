package math

import (
	"go-jvm/src/ch11/instruction/base"
	"go-jvm/src/ch11/rtda"
	"go-jvm/src/ch11/utils"
)

type IADD struct {
	base.NoOperandsInstruction
}

func (i *IADD) Execute(frame *rtda.Frame) {
	val1 := int64(frame.OperandStack.PopInt())
	val2 := int64(frame.OperandStack.PopInt())
	frame.OperandStack.PushInt(int32(val1 + val2))
}

type ISUB struct {
	base.NoOperandsInstruction
}

func (i *ISUB) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopInt()
	val2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val2 - val1)
}

type IMUL struct {
	base.NoOperandsInstruction
}

func (i *IMUL) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopInt()
	val2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val1 * val2)
}

type IDIV struct {
	base.NoOperandsInstruction
}

func (i *IDIV) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopInt()
	val2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val2 / val1)
}

type INEG struct {
	base.NoOperandsInstruction
}

func (i *INEG) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(-frame.OperandStack.PopInt())
}

type IREM struct {
	base.NoOperandsInstruction
}

func (i *IREM) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopInt()
	val2 := frame.OperandStack.PopInt()
	if val1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	frame.OperandStack.PushInt(val2 % val1)
}

type IAND struct {
	base.NoOperandsInstruction
}

func (i *IAND) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopInt()
	val2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val1 & val2)
}

type IOR struct {
	base.NoOperandsInstruction
}

func (i *IOR) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopInt()
	val2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val1 | val2)
}

type IXOR struct {
	base.NoOperandsInstruction
}

func (i *IXOR) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopInt()
	val2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val1 ^ val2)
}

type ISHL struct {
	base.NoOperandsInstruction
}

func (i *ISHL) Execute(frame *rtda.Frame) {
	val1 := uint32(frame.OperandStack.PopInt()) & 31
	val2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val2 << val1)
}

type ISHR struct {
	base.NoOperandsInstruction
}

func (i *ISHR) Execute(frame *rtda.Frame) {
	val1 := uint32(frame.OperandStack.PopInt()) & 31
	val2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val2 >> val1)
}

type IUSHL struct {
	base.NoOperandsInstruction
}

func (i *IUSHL) Execute(frame *rtda.Frame) {
	val1 := uint32(frame.OperandStack.PopInt()) & 31
	val2 := uint32(frame.OperandStack.PopInt())
	frame.OperandStack.PushInt(int32(val2 << val1))
}

type IUSHR struct {
	base.NoOperandsInstruction
}

func (i *IUSHR) Execute(frame *rtda.Frame) {
	val1 := uint32(frame.OperandStack.PopInt()) & 31
	val2 := uint32(frame.OperandStack.PopInt())
	frame.OperandStack.PushInt(int32(val2 >> val1))
}

type IINC struct {
	Index uint
	Const int32
}

func (i *IINC) FetchOperands(cr *utils.ByteCodeReader) {
	i.Index = uint(cr.ReadUint8())
	i.Const = int32(int8(cr.ReadUint8()))
}

func (i *IINC) Execute(frame *rtda.Frame) {
	val := frame.LocalVars.GetInt(i.Index)
	frame.LocalVars.SetInt(i.Index, val+i.Const)
}
