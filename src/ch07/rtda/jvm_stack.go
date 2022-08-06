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

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if s.tail == nil {
		s.tail = frame
	} else {
		frame.prev = s.tail
		s.tail = frame
	}
	s.size++
}

func (s *Stack) pop() *Frame {
	if s.size < 0 {
		panic("Empty jvm stack")
	}
	frame := s.tail
	s.tail = s.tail.prev
	frame.prev = nil
	s.size--
	return frame
}

func (s *Stack) peek() *Frame {
	if s.size < 0 {
		return nil
	}
	return s.tail
}
