package control

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"go-jvm/src/utils"
)

type GOTO struct {
	base.BranchInstruction
}

func (inst *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, inst.Offset)
}

type GOTO_W struct {
	Offset int
}

func (inst *GOTO_W) FetchOperands(reader *utils.ByteCodeReader) {
	inst.Offset = int(reader.ReadUint32())
}

func (inst *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, inst.Offset)
}
