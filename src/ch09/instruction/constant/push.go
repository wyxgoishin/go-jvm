package constant

import (
	"go-jvm/src/ch09/rtda"
	"go-jvm/src/ch09/utils"
)

type BI_PUSH struct {
	val int8 // push byte
}

func (inst *BI_PUSH) FetchOperands(cr *utils.ByteCodeReader) {
	inst.val = int8(cr.ReadUint8())
}

func (inst *BI_PUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(int32(inst.val))
}

type SI_PUSH struct {
	val int16 // push short
}

func (inst *SI_PUSH) FetchOperands(cr *utils.ByteCodeReader) {
	inst.val = int16(cr.ReadUint16())
}

func (inst *SI_PUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(int32(inst.val))
}
