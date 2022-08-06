package constant

import (
	"go-jvm/src/ch07/rtda"
	"go-jvm/src/ch07/utils"
)

type BI_PUSH struct {
	val int8 // push byte
}

func (b *BI_PUSH) FetchOperands(cr *utils.ByteCodeReader) {
	b.val = int8(cr.ReadUint8())
}

func (b *BI_PUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(int32(b.val))
}

type SI_PUSH struct {
	val int16 // push short
}

func (s *SI_PUSH) FetchOperands(cr *utils.ByteCodeReader) {
	s.val = int16(cr.ReadUint16())
}

func (s *SI_PUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(int32(s.val))
}
