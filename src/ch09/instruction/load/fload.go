package load

import (
	"go-jvm/src/ch09/instruction/base"
	"go-jvm/src/ch09/rtda"
)

type FLOAD struct {
	base.Index8Struction
}

func (f *FLOAD) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushFloat(frame.LocalVars.GetFloat(f.Index))
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_0) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushFloat(frame.LocalVars.GetFloat(0))
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_1) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushFloat(frame.LocalVars.GetFloat(1))
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_2) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushFloat(frame.LocalVars.GetFloat(2))
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_3) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushFloat(frame.LocalVars.GetFloat(3))
}
