package constant

import (
	"go-jvm/src/ch10/instruction/base"
	"go-jvm/src/ch10/rtda"
	"go-jvm/src/ch10/rtda/heap"
)

type LDC struct {
	base.Index8Struction
}

func (inst *LDC) Execute(frame *rtda.Frame) {
	ldc(frame, inst.Index)
}

type LDC_W struct {
	base.Index16Struction
}

func (inst *LDC_W) Execute(frame *rtda.Frame) {
	ldc(frame, inst.Index)
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
	case string:
		internedString := heap.JvString(frame.Method().Class().ClassLoader(), con.(string))
		stack.PushRef(internedString)
	case *heap.ClassRef:
		classRef := con.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JvClass()
		stack.PushRef(classObj)
	default:
		panic("todo: ldc")
	}
}

type LDC2_W struct {
	base.Index16Struction
}

func (inst *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack
	rtcp := frame.Method().Class().RuntimeConstantPool()
	con := rtcp.GetConstant(inst.Index)
	switch con.(type) {
	case int64:
		stack.PushLong(con.(int64))
	case float64:
		stack.PushDouble(con.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
