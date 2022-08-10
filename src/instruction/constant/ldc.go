package constant

import (
	"go-jvm/src/instruction/base"
	"go-jvm/src/rtda"
	"go-jvm/src/rtda/heap"
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
	stack := frame.OperandStack()
	rtcp := frame.Method().Class().RuntimeConstantPool()
	cval := rtcp.GetConstant(index)
	switch cval.(type) {
	case int32:
		stack.PushInt(cval.(int32))
	case float32:
		stack.PushFloat(cval.(float32))
	case string:
		internedString := heap.JvString(frame.Method().Class().ClassLoader(), cval.(string))
		stack.PushRef(internedString)
	case *heap.ClassRef:
		classRef := cval.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JvClass()
		stack.PushRef(classObj)
	default:
		//ToDo: ldc
		panic("todo: ldc")
	}
}

type LDC2_W struct {
	base.Index16Struction
}

func (inst *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	rtcp := frame.Method().Class().RuntimeConstantPool()
	cval := rtcp.GetConstant(inst.Index)
	switch cval.(type) {
	case int64:
		stack.PushLong(cval.(int64))
	case float64:
		stack.PushDouble(cval.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
