package load

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type LLOAD struct {
	base.Index8Struction
}

func (inst *LLOAD) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(frame.LocalVars().GetLong(inst.Index))
}

type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (inst *LLOAD_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(frame.LocalVars().GetLong(0))
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (inst *LLOAD_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(frame.LocalVars().GetLong(1))
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (inst *LLOAD_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(frame.LocalVars().GetLong(2))
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (inst *LLOAD_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(frame.LocalVars().GetLong(3))
}
