package control

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
	"go-jvm/src/ch06/utils"
)

type GOTO struct {
	base.BranchInstruction
}

func (g *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, g.Offset)
}

type GOTO_W struct {
	Offset int
}

func (g *GOTO_W) FetchOperands(cr *utils.ByteCodeReader) {
	g.Offset = int(int32(cr.ReadUint32()))
}

func (g *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, g.Offset)
}
