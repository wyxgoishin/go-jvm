package math

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"go-jvm/src/utils"
)

type IADD struct {
	base.NoOperandsInstruction
}

func (inst *IADD) Execute(frame *rtda.Frame) {
	val2 := int64(frame.OperandStack().PopInt())
	val1 := int64(frame.OperandStack().PopInt())
	frame.OperandStack().PushInt(int32(val1 + val2))
}

type ISUB struct {
	base.NoOperandsInstruction
}

func (inst *ISUB) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(val1 - val2)
}

type IMUL struct {
	base.NoOperandsInstruction
}

func (inst *IMUL) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(val1 * val2)
}

type IDIV struct {
	base.NoOperandsInstruction
}

func (inst *IDIV) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(val1 / val2)
}

type INEG struct {
	base.NoOperandsInstruction
}

func (inst *INEG) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(-val)
}

type IREM struct {
	base.NoOperandsInstruction
}

//ToDo: support throw Exception
func (inst *IREM) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	frame.OperandStack().PushInt(val1 % val2)
}

type IAND struct {
	base.NoOperandsInstruction
}

func (inst *IAND) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(val1 & val2)
}

type IOR struct {
	base.NoOperandsInstruction
}

func (inst *IOR) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(val1 | val2)
}

type IXOR struct {
	base.NoOperandsInstruction
}

func (inst *IXOR) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(val1 ^ val2)
}

type ISHL struct {
	base.NoOperandsInstruction
}

func (inst *ISHL) Execute(frame *rtda.Frame) {
	val2 := uint32(frame.OperandStack().PopInt()) & 31
	val1 := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(val1 << val2)
}

type ISHR struct {
	base.NoOperandsInstruction
}

func (inst *ISHR) Execute(frame *rtda.Frame) {
	val2 := uint32(frame.OperandStack().PopInt()) & 31
	val1 := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(val1 >> val2)
}

type IUSHL struct {
	base.NoOperandsInstruction
}

func (inst *IUSHL) Execute(frame *rtda.Frame) {
	val2 := uint32(frame.OperandStack().PopInt()) & 31
	val1 := uint32(frame.OperandStack().PopInt())
	frame.OperandStack().PushInt(int32(val1 << val2))
}

type IUSHR struct {
	base.NoOperandsInstruction
}

func (inst *IUSHR) Execute(frame *rtda.Frame) {
	val2 := uint32(frame.OperandStack().PopInt()) & 31
	val1 := uint32(frame.OperandStack().PopInt())
	frame.OperandStack().PushInt(int32(val1 >> val2))
}

type IINC struct {
	Index uint
	Const int32
}

func (inst *IINC) FetchOperands(reader *utils.ByteCodeReader) {
	inst.Index = uint(reader.ReadUint8())
	inst.Const = int32(reader.ReadInt8())
}

func (inst *IINC) Execute(frame *rtda.Frame) {
	val := frame.LocalVars().GetInt(inst.Index)
	frame.LocalVars().SetInt(inst.Index, val+inst.Const)
}
