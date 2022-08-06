package constant

import (
	"go-jvm/src/ch07/instruction/base"
	"go-jvm/src/ch07/rtda"
)

type LDC struct {
	base.Index8Struction
}

func (l *LDC) Execute(frame *rtda.Frame) {
	ldc(frame, l.Index)
}

type LDC_W struct {
	base.Index16Struction
}

func (l *LDC_W) Execute(frame *rtda.Frame) {
	ldc(frame, l.Index)
}

func ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack
	rtcp := frame.Method().Class().RuntimeConstantPool()
	con := rtcp.GetConstant(index)
	switch con.(type) {
	case int32:
		stack.PushInt(con.(int32))
	case float32:
		stack.PushFloat(con.(float32))
	//case string:
	// case *heap.ClassRef:
	default:
		panic("todo: ldc")
	}
}

type LDC2_W struct {
	base.Index16Struction
}

func (l *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	rtcp := frame.Method().Class().RuntimeConstantPool()
	con := rtcp.GetConstant(l.Index)
	switch con.(type) {
	case int64:
		stack.PushLong(con.(int64))
	case float32:
		stack.PushDouble(con.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
