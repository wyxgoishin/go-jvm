package rtda

type Frame struct {
	prev         *Frame
	LocalVars    LocalVars
	OperandStack *OperandStack
	thread       *Thread
	nextPC       int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		LocalVars:    newLocalVars(maxLocals),
		OperandStack: newOperandStack(maxStack),
		thread:       thread,
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
