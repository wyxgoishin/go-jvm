package rtda

import (
	"fmt"
	"go-jvm/src/rtda/heap"
	"math"
)

type OperandStack struct {
	size   uint
	slots  []Slot
	valIdx map[uint]bool
}

func NewOperandStack(maxStack uint) *OperandStack {
	if maxStack <= 0 {
		return nil
	}
	return &OperandStack{
		slots:  make([]Slot, maxStack),
		valIdx: map[uint]bool{},
	}
}

func (stack *OperandStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *OperandStack) PushInt(val int32) {
	stack.slots[stack.size].num = val
	stack.valIdx[stack.size] = true
	stack.size++
}

func (stack *OperandStack) PopInt() int32 {
	if stack.size == 0 {
		panic("Empty stack frame operandStack")
	}
	stack.size--
	stack.valIdx[stack.size] = false
	return stack.slots[stack.size].num
}

func (stack *OperandStack) PushFloat(val float32) {
	ival := int32(math.Float32bits(val))
	stack.PushInt(ival)
}

func (stack *OperandStack) PopFloat() float32 {
	return math.Float32frombits(uint32(stack.PopInt()))
}

func (stack *OperandStack) PushLong(val int64) {
	stack.slots[stack.size].num = int32(val)
	stack.valIdx[stack.size] = true
	stack.size++
	stack.slots[stack.size].num = int32(val >> 32)
	stack.size++
}

func (stack *OperandStack) PopLong() int64 {
	if stack.size == 0 {
		panic("Empty stack frame operandStack")
	}
	higher := int64(stack.PopInt()) << 32
	lower := int64(stack.PopInt())
	return higher + lower
}

func (stack *OperandStack) PushDouble(val float64) {
	ival := int64(math.Float64bits(val))
	stack.PushLong(ival)
}

func (stack *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(stack.PopLong()))
}

func (stack *OperandStack) PushRef(val *heap.Object) {
	stack.slots[stack.size].ref = val
	stack.valIdx[stack.size] = true
	stack.size++
}

func (stack *OperandStack) PopRef() *heap.Object {
	if stack.size == 0 {
		panic("Empty stack frame operandStack")
	}
	stack.size--
	stack.valIdx[stack.size] = false
	ref := stack.slots[stack.size].ref
	// 帮助Go的垃圾收集器回收Object结构体实例
	stack.slots[stack.size].ref = nil
	return ref
}

func (stack *OperandStack) PushSlot(slot Slot) {
	stack.slots[stack.size] = slot
	stack.valIdx[stack.size] = true
	stack.size++
}

func (stack *OperandStack) PopSlot() Slot {
	if stack.size == 0 {
		panic("Empty stack frame operandStack")
	}
	stack.size--
	stack.valIdx[stack.size] = false
	return stack.slots[stack.size]
}

func (stack *OperandStack) PushTwoSlot(slot1, slot2 Slot) {
	stack.slots[stack.size] = slot1
	stack.valIdx[stack.size] = true
	stack.size++
	stack.slots[stack.size] = slot2
	stack.size++
}

func (stack *OperandStack) IsTopSlotStartOfVal() bool {
	return stack.valIdx[stack.size-1]
}

func (stack *OperandStack) GetRefFromTop(idx uint) *heap.Object {
	if stack.size-1-idx < 0 {
		panic(fmt.Sprintf("Invalid idx: %v", idx))
	}
	return stack.slots[stack.size-1-idx].ref
}

func (stack *OperandStack) PushBoolean(bval bool) {
	var ival int32
	if bval {
		ival = 1
	}
	stack.PushInt(ival)
}

func (stack *OperandStack) PopBoolean() bool {
	ival := stack.PopInt()
	if ival == 0 {
		return false
	}
	return true
}

func (stack *OperandStack) Clear() {
	stack.size = 0
	for key := range stack.valIdx {
		stack.valIdx[key] = false
		stack.slots[key].ref = nil
	}
}
