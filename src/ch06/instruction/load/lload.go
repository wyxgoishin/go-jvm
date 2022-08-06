package load

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
)

type LLOAD struct {
	base.Index8Struction
}

func (l *LLOAD) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(l.Index))
}

type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_0) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(0))
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_1) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(1))
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_2) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(2))
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_3) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(3))
}
