package rtda

import "go-jvm/src/ch10/rtda/heap"

type Frame struct {
	prev         *Frame
	LocalVars    LocalVars
	OperandStack *OperandStack
	thread       *Thread
	nextPC       int
	method       *heap.Method
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		LocalVars:    newLocalVars(method.MaxLocals()),
		OperandStack: newOperandStack(method.MaxStack()),
		thread:       thread,
		method:       method,
	}
}

func (frame *Frame) Thread() *Thread {
	return frame.thread
}

func (frame *Frame) NextPC() int {
	return frame.nextPC
}

func (frame *Frame) SetNextPC(pc int) {
	frame.nextPC = pc
}

func (frame *Frame) Method() *heap.Method {
	return frame.method
}

func (frame *Frame) RevertNextPC() {
	frame.nextPC = frame.thread.pc
}
