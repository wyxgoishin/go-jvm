package comparison

import (
	"go-jvm/src/ch10/instruction/base"
	"go-jvm/src/ch10/rtda"
)

type IF_EQ struct {
	base.BranchInstruction
}

func (i *IF_EQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val == 0 {
		base.Branch(frame, i.Offset)
	}
}

type IF_NE struct {
	base.BranchInstruction
}

func (i *IF_NE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val != 0 {
		base.Branch(frame, i.Offset)
	}
}

type IF_LT struct {
	base.BranchInstruction
}

func (i *IF_LT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val < 0 {
		base.Branch(frame, i.Offset)
	}
}

type IF_LE struct {
	base.BranchInstruction
}

func (i *IF_LE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val <= 0 {
		base.Branch(frame, i.Offset)
	}
}

type IF_GT struct {
	base.BranchInstruction
}

func (i *IF_GT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val > 0 {
		base.Branch(frame, i.Offset)
	}
}

type IF_GE struct {
	base.BranchInstruction
}

func (i *IF_GE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	if val >= 0 {
		base.Branch(frame, i.Offset)
	}
}

type IF_NON_NULL struct {
	base.BranchInstruction
}

func (i *IF_NON_NULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack.PopRef()
	if ref != nil {
		base.Branch(frame, i.Offset)
	}
}

type IF_NULL struct {
	base.BranchInstruction
}

func (i *IF_NULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack.PopRef()
	if ref == nil {
		base.Branch(frame, i.Offset)
	}
}

type IF_ACMP_EQ struct {
	base.BranchInstruction
}

func (i *IF_ACMP_EQ) Execute(frame *rtda.Frame) {
	ref1 := frame.OperandStack.PopRef()
	ref2 := frame.OperandStack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, i.Offset)
	}
}

type IF_ACMP_NE struct {
	base.BranchInstruction
}

func (i *IF_ACMP_NE) Execute(frame *rtda.Frame) {
	ref1 := frame.OperandStack.PopRef()
	ref2 := frame.OperandStack.PopRef()
	if ref1 != ref2 {
		base.Branch(frame, i.Offset)
	}
}

type IF_ICMP_EQ struct {
	base.BranchInstruction
}

func (i *IF_ICMP_EQ) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopInt()
	val2 := frame.OperandStack.PopInt()
	if val1 == val2 {
		base.Branch(frame, i.Offset)
	}
}

type IF_ICMP_NE struct {
	base.BranchInstruction
}

func (i *IF_ICMP_NE) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack.PopInt()
	val2 := frame.OperandStack.PopInt()
	if val1 != val2 {
		base.Branch(frame, i.Offset)
	}
}

type IF_ICMP_LT struct {
	base.BranchInstruction
}

func (i *IF_ICMP_LT) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack.PopInt()
	val1 := frame.OperandStack.PopInt()
	if val1 < val2 {
		base.Branch(frame, i.Offset)
	}
}

type IF_ICMP_LE struct {
	base.BranchInstruction
}

func (i *IF_ICMP_LE) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack.PopInt()
	val1 := frame.OperandStack.PopInt()
	if val1 <= val2 {
		base.Branch(frame, i.Offset)
	}
}

type IF_ICMP_GT struct {
	base.BranchInstruction
}

func (i *IF_ICMP_GT) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack.PopInt()
	val1 := frame.OperandStack.PopInt()
	if val1 > val2 {
		base.Branch(frame, i.Offset)
	}
}

type IF_ICMP_GE struct {
	base.BranchInstruction
}

func (i *IF_ICMP_GE) Execute(frame *rtda.Frame) {
	val2 := frame.OperandStack.PopInt()
	val1 := frame.OperandStack.PopInt()
	if val1 >= val2 {
		base.Branch(frame, i.Offset)
	}
}
