package load

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type DLOAD struct {
	base.Index8Struction
}

func (inst *DLOAD) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(frame.LocalVars().GetDouble(inst.Index))
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (inst *DLOAD_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(frame.LocalVars().GetDouble(0))
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (inst *DLOAD_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(frame.LocalVars().GetDouble(1))
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (inst *DLOAD_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(frame.LocalVars().GetDouble(2))
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (inst *DLOAD_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(frame.LocalVars().GetDouble(3))
}
