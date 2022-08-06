package stack

import (
	"go-jvm/src/ch06/instruction/base"
	"go-jvm/src/ch06/rtda"
)

type DUP struct {
	base.NoOperandsInstruction
}

func (d *DUP) Execute(frame *rtda.Frame) {
	slot := frame.OperandStack.PopSlot()
	frame.OperandStack.PushSlot(slot)
	frame.OperandStack.PushSlot(slot)
}

type DUP_X1 struct {
	base.NoOperandsInstruction
}

func (d *DUP_X1) Execute(frame *rtda.Frame) {
	slot1 := frame.OperandStack.PopSlot()
	slot2 := frame.OperandStack.PopSlot()
	frame.OperandStack.PushSlot(slot1)
	frame.OperandStack.PushSlot(slot2)
	frame.OperandStack.PushSlot(slot1)
}

type DUP_X2 struct {
	base.NoOperandsInstruction
}

func (d *DUP_X2) Execute(frame *rtda.Frame) {
	slot1 := frame.OperandStack.PopSlot()
	isVal2TwoSlot := frame.OperandStack.IsTopSlotStartOfVal()
	slot2 := frame.OperandStack.PopSlot()
	slot3 := frame.OperandStack.PopSlot()
	frame.OperandStack.PushSlot(slot1)
	if isVal2TwoSlot {
		frame.OperandStack.PushTwoSlot(slot3, slot2)
	} else {
		frame.OperandStack.PushSlot(slot3)
		frame.OperandStack.PushSlot(slot2)
	}
	frame.OperandStack.PushSlot(slot1)
}

type DUP2 struct {
	base.NoOperandsInstruction
}

func (d *DUP2) Execute(frame *rtda.Frame) {
	isVal1TwoSlot := frame.OperandStack.IsTopSlotStartOfVal()
	slot1 := frame.OperandStack.PopSlot()
	slot2 := frame.OperandStack.PopSlot()
	if isVal1TwoSlot {
		frame.OperandStack.PushTwoSlot(slot2, slot1)
		frame.OperandStack.PushTwoSlot(slot2, slot1)
	} else {
		frame.OperandStack.PushSlot(slot2)
		frame.OperandStack.PushSlot(slot1)
		frame.OperandStack.PushSlot(slot2)
		frame.OperandStack.PushSlot(slot1)
	}
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

func (d *DUP2_X1) Execute(frame *rtda.Frame) {
	isVal1TwoSlot := frame.OperandStack.IsTopSlotStartOfVal()
	slot1 := frame.OperandStack.PopSlot()
	slot2 := frame.OperandStack.PopSlot()
	slot3 := frame.OperandStack.PopSlot()
	if isVal1TwoSlot {
		frame.OperandStack.PushSlot(slot2)
		frame.OperandStack.PushSlot(slot1)
		frame.OperandStack.PushSlot(slot3)
		frame.OperandStack.PushSlot(slot2)
		frame.OperandStack.PushSlot(slot1)
	} else {
		frame.OperandStack.PushTwoSlot(slot2, slot1)
		frame.OperandStack.PushSlot(slot3)
		frame.OperandStack.PushTwoSlot(slot2, slot1)
	}
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (d *DUP2_X2) Execute(frame *rtda.Frame) {
	var slot1, slot2, slot3, slot4 rtda.Slot
	var isVal1TwoSlot, isVal2TwoSlot, isVal3TwoSlot bool

	isVal1TwoSlot = frame.OperandStack.IsTopSlotStartOfVal()
	slot1 = frame.OperandStack.PopSlot()
	if isVal1TwoSlot {
		slot2 = frame.OperandStack.PopSlot()
		isVal2TwoSlot = frame.OperandStack.IsTopSlotStartOfVal()
	} else {
		isVal2TwoSlot = frame.OperandStack.IsTopSlotStartOfVal()
		slot2 = frame.OperandStack.PopSlot()
	}

	isVal3TwoSlot = frame.OperandStack.IsTopSlotStartOfVal()
	slot3 = frame.OperandStack.PopSlot()
	slot4 = frame.OperandStack.PopSlot()

	if isVal2TwoSlot && isVal1TwoSlot {
		frame.OperandStack.PushTwoSlot(slot2, slot1)
		frame.OperandStack.PushTwoSlot(slot4, slot3)
		frame.OperandStack.PushTwoSlot(slot2, slot1)
	} else if isVal3TwoSlot {
		frame.OperandStack.PushSlot(slot2)
		frame.OperandStack.PushSlot(slot1)
		frame.OperandStack.PushTwoSlot(slot4, slot3)
		frame.OperandStack.PushSlot(slot2)
		frame.OperandStack.PushSlot(slot1)
	} else if isVal1TwoSlot {
		frame.OperandStack.PushTwoSlot(slot2, slot1)
		frame.OperandStack.PushSlot(slot4)
		frame.OperandStack.PushSlot(slot3)
		frame.OperandStack.PushTwoSlot(slot2, slot1)
	}
}
