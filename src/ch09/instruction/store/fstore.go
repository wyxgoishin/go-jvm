package store

import (
	"go-jvm/src/ch09/instruction/base"
	"go-jvm/src/ch09/rtda"
)

type FSTORE struct {
	base.Index8Struction
}

func (f *FSTORE) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetFloat(f.Index, frame.OperandStack.PopFloat())
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (f *FSTORE_0) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetFloat(0, frame.OperandStack.PopFloat())
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (f *FSTORE_1) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetFloat(1, frame.OperandStack.PopFloat())
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (f *FSTORE_2) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetFloat(2, frame.OperandStack.PopFloat())
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (f *FSTORE_3) Execute(frame *rtda.Frame) {
	frame.LocalVars.SetFloat(3, frame.OperandStack.PopFloat())
}
