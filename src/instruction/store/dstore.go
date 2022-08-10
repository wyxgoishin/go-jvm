package store

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type DSTORE struct {
	base.Index8Struction
}

func (inst *DSTORE) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetDouble(inst.Index, frame.OperandStack().PopDouble())
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (inst *DSTORE_0) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetDouble(0, frame.OperandStack().PopDouble())
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (inst *DSTORE_1) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetDouble(1, frame.OperandStack().PopDouble())
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (inst *DSTORE_2) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetDouble(2, frame.OperandStack().PopDouble())
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (inst *DSTORE_3) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetDouble(3, frame.OperandStack().PopDouble())
}
