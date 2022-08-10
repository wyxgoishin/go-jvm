package constant

import (
	"go-jvm/src/rtda"
	"go-jvm/src/utils"
)

type BI_PUSH struct {
	val int8 // push byte
}

func (inst *BI_PUSH) FetchOperands(reader *utils.ByteCodeReader) {
	inst.val = reader.ReadInt8()
}

func (inst *BI_PUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(inst.val))
}

type SI_PUSH struct {
	val int16 // push short
}

func (inst *SI_PUSH) FetchOperands(reader *utils.ByteCodeReader) {
	inst.val = reader.ReadInt16()
}

func (inst *SI_PUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(inst.val))
}
