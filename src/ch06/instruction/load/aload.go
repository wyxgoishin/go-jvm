package load

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
)

type ALOAD struct {
	base.Index8Struction
}

func (a *ALOAD) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(a.Index))
}

type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (a *ALOAD_0) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(0))
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (a *ALOAD_1) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(1))
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (a *ALOAD_2) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(2))
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (a *ALOAD_3) Execute(frame *rtda.Frame) {
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(3))
}
