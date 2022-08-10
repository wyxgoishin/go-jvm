package load

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type FLOAD struct {
	base.Index8Struction
}

func (inst *FLOAD) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(frame.LocalVars().GetFloat(inst.Index))
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (inst *FLOAD_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(frame.LocalVars().GetFloat(0))
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (inst *FLOAD_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(frame.LocalVars().GetFloat(1))
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (inst *FLOAD_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(frame.LocalVars().GetFloat(2))
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (inst *FLOAD_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(frame.LocalVars().GetFloat(3))
}
