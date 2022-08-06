package rtda

import "go-jvm/src/ch10/rtda/heap"

type Thread struct {
	pc    int
	stack *Stack
}

const (
	StackSize = 1024
)

func NewThread() *Thread {
	return &Thread{
		stack: newStack(StackSize),
	}
}

func (thread *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(thread, method)
}

func (thread *Thread) PC() int {
	return thread.pc
}

func (thread *Thread) SetPC(pc int) {
	thread.pc = pc
}

func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}

func (thread *Thread) PopFrame() *Frame {
	return thread.stack.pop()
}

func (thread *Thread) CurrentFrame() *Frame {
	return thread.stack.peek()
}

func (thread *Thread) TopFrame() *Frame {
	return thread.stack.peek()
}

func (thread *Thread) IsStackEmpty() bool {
	return thread.stack.IsEmpty()
}

func (thread *Thread) ClearStack() {
	for !thread.IsStackEmpty() {
		thread.PopFrame()
	}
}

func (thread *Thread) GetFrames() []*Frame {
	return thread.stack.getFrames()
}
