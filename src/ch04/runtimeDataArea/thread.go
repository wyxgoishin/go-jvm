package runtimeDataArea

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

func (t *Thread) PC() int {
	return t.pc
}

func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.peek()
}
