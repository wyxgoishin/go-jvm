package store

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type ISTORE struct {
	base.Index8Struction
}

func (inst *ISTORE) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetInt(inst.Index, frame.OperandStack().PopInt())
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (inst *ISTORE_0) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetInt(0, frame.OperandStack().PopInt())
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (inst *ISTORE_1) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetInt(1, frame.OperandStack().PopInt())
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (inst *ISTORE_2) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetInt(2, frame.OperandStack().PopInt())
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (inst *ISTORE_3) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetInt(3, frame.OperandStack().PopInt())
}
