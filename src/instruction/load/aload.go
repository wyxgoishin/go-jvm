package load

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type ALOAD struct {
	base.Index8Struction
}

func (inst *ALOAD) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(frame.LocalVars().GetRef(inst.Index))
}

type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (inst *ALOAD_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(frame.LocalVars().GetRef(0))
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (inst *ALOAD_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(frame.LocalVars().GetRef(1))
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (inst *ALOAD_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(frame.LocalVars().GetRef(2))
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (inst *ALOAD_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(frame.LocalVars().GetRef(3))
}
