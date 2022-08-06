package store

import (
	"go-jvm/src/ch10/instruction/base"
	"go-jvm/src/ch10/rtda"
)

type LSTORE struct {
	base.Index8Struction
}

func (l *LSTORE) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetLong(l.Index, frame.OperandStack.PopLong())
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (l *LSTORE_0) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetLong(0, frame.OperandStack.PopLong())
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (l *LSTORE_1) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetLong(1, frame.OperandStack.PopLong())
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (l *LSTORE_2) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetLong(2, frame.OperandStack.PopLong())
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (l *LSTORE_3) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetLong(3, frame.OperandStack.PopLong())
}
