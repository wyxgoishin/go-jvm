package load

import (
	"go-jvm/src/ch11/instruction/base"
	"go-jvm/src/ch11/rtda"
)

type ILOAD struct {
	base.Index8Struction
}

func (i *ILOAD) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(frame.LocalVars.GetInt(i.Index))
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_0) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(frame.LocalVars.GetInt(0))
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_1) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(frame.LocalVars.GetInt(1))
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_2) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(frame.LocalVars.GetInt(2))
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_3) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushInt(frame.LocalVars.GetInt(3))
}
