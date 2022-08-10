package control

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction
}

func (inst *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type ARETURN struct {
	base.NoOperandsInstruction
}

func (inst *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	curFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := curFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

type DRETURN struct {
	base.NoOperandsInstruction
}

func (inst *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	curFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := curFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}

type FRETURN struct {
	base.NoOperandsInstruction
}

func (f *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	curFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := curFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}

type IRETURN struct {
	base.NoOperandsInstruction
}

func (inst *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	curFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := curFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

type LRETURN struct {
	base.NoOperandsInstruction
}

func (inst *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	curFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := curFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}
