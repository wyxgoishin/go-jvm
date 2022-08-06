package store

import (
	"go-jvm/src/ch07/instruction/base"
	"go-jvm/src/ch07/rtda"
)

type ASTORE struct {
	base.Index8Struction
}

func (a *ASTORE) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetRef(a.Index, frame.OperandStack.PopRef())
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (a *ASTORE_0) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetRef(0, frame.OperandStack.PopRef())
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (a *ASTORE_1) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetRef(1, frame.OperandStack.PopRef())
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (a *ASTORE_2) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetRef(2, frame.OperandStack.PopRef())
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (a *ASTORE_3) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetRef(3, frame.OperandStack.PopRef())
}
