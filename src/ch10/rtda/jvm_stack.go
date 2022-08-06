package rtda

type Stack struct {
	maxSize uint
	size    uint
	tail    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (stack *Stack) Size() uint {
	return stack.size
}

func (stack *Stack) getFrames() []*Frame {
	frames := make([]*Frame, 0)
	for cur := stack.tail; cur != nil; cur = cur.prev {
		frames = append(frames, cur)
	}
	return frames
}

func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *Stack) push(frame *Frame) {
	if stack.size >= stack.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if stack.tail == nil {
		stack.tail = frame
	} else {
		frame.prev = stack.tail
		stack.tail = frame
	}
	stack.size++
}

func (stack *Stack) pop() *Frame {
	if stack.size < 0 {
		panic("Empty jvm stack")
	}
	frame := stack.tail
	stack.tail = stack.tail.prev
	frame.prev = nil
	stack.size--
	return frame
}

func (stack *Stack) peek() *Frame {
	if stack.size < 0 {
		return nil
	}
	return stack.tail
}
