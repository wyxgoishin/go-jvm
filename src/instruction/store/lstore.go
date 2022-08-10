package store

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type LSTORE struct {
	base.Index8Struction
}

func (ins *LSTORE) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetLong(ins.Index, frame.OperandStack().PopLong())
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (ins *LSTORE_0) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetLong(0, frame.OperandStack().PopLong())
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (ins *LSTORE_1) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetLong(1, frame.OperandStack().PopLong())
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (ins *LSTORE_2) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetLong(2, frame.OperandStack().PopLong())
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (ins *LSTORE_3) Execute(frame *rtda.Frame) {
	frame.LocalVars().SetLong(3, frame.OperandStack().PopLong())
}
