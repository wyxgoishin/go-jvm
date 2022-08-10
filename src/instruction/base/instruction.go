package base

import (
	"go-jvm/src/rtda"
	"go-jvm/src/utils"
)

type Instruction interface {
	FetchOperands(reader *utils.ByteCodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (inst *NoOperandsInstruction) FetchOperands(reader *utils.ByteCodeReader) {
}

type Index8Struction struct {
	Index uint
}

func (inst *Index8Struction) FetchOperands(reader *utils.ByteCodeReader) {
	inst.Index = uint(reader.ReadUint8())
}

type Index16Struction struct {
	Index uint
}

func (inst *Index16Struction) FetchOperands(reader *utils.ByteCodeReader) {
	inst.Index = uint(reader.ReadUint16())
}

type BranchInstruction struct {
	Offset int
}

func (inst *BranchInstruction) FetchOperands(reader *utils.ByteCodeReader) {
	inst.Offset = int(int16(reader.ReadUint16()))
}

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
