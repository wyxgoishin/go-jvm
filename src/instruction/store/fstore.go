package store

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type FSTORE struct {
	base.Index8Struction
}

func (inst *FSTORE) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetFloat(inst.Index, frame.OperandStack().PopFloat())
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (inst *FSTORE_0) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetFloat(0, frame.OperandStack().PopFloat())
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (inst *FSTORE_1) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetFloat(1, frame.OperandStack().PopFloat())
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (inst *FSTORE_2) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetFloat(2, frame.OperandStack().PopFloat())
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (inst *FSTORE_3) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetFloat(3, frame.OperandStack().PopFloat())
}
