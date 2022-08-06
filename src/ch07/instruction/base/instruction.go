package base

import (
	"go-jvm/src/ch07/rtda"
	"go-jvm/src/ch07/utils"
)

type Instruction interface {
	FetchOperands(cr *utils.ByteCodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (n *NoOperandsInstruction) FetchOperands(cr *utils.ByteCodeReader) {
}

func (n *NoOperandsInstruction) Execute(frame *rtda.Frame) {

}

type BranchInstruction struct {
	Offset int
}

func (b *BranchInstruction) FetchOperands(cr *utils.ByteCodeReader) {
	b.Offset = int(int16(cr.ReadUint16()))
}

func (b *BranchInstruction) Execute(frame *rtda.Frame) {

}

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}

type Index8Struction struct {
	Index uint
}

func (i *Index8Struction) FetchOperands(cr *utils.ByteCodeReader) {
	i.Index = uint(cr.ReadUint8())
}

type Index16Struction struct {
	Index uint
}

func (i *Index16Struction) FetchOperands(cr *utils.ByteCodeReader) {
	i.Index = uint(cr.ReadUint16())
}
