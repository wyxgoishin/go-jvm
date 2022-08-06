package rtda

import "go-jvm/src/ch07/rtda/heap"

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

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) NextPC() int {
	return f.nextPC
}

func (f *Frame) SetNextPC(pc int) {
	f.nextPC = pc
}

func (f *Frame) Method() *heap.Method {
	return f.method
}

func (f *Frame) RevertNextPC() {
	f.nextPC = f.thread.pc
}
