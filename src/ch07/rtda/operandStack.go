package rtda

import (
	"go-jvm/src/ch07/rtda/heap"
	"math"
)

type OperandStack struct {
	size   uint
	slots  []Slot
	valIdx map[uint]bool
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack <= 0 {
		return nil
	}
	return &OperandStack{
		slots:  make([]Slot, maxStack),
		valIdx: map[uint]bool{},
	}
}

func (o *OperandStack) IsEmpty() bool {
	return o.size == 0
}

func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].num = val
	o.valIdx[o.size] = true
	o.size++
}

func (o *OperandStack) PopInt() int32 {
	o.size--
	o.valIdx[o.size] = false
	return o.slots[o.size].num
}

func (o *OperandStack) PushFloat(val float32) {
	ival := int32(math.Float32bits(val))
	o.PushInt(ival)
}

func (o *OperandStack) PopFloat() float32 {
	o.size--
	return math.Float32frombits(uint32(o.slots[o.size].num))
}

func (o *OperandStack) PushLong(val int64) {
	o.slots[o.size].num = int32(val)
	o.valIdx[o.size] = true
	o.size++
	o.slots[o.size].num = int32(val >> 32)
	o.size++
}

func (o *OperandStack) PopLong() int64 {
	higher := int64(o.PopInt()) << 32
	lower := int64(o.PopInt())
	return higher + lower
}

func (o *OperandStack) PushDouble(val float64) {
	ival := int64(math.Float64bits(val))
	o.PushLong(ival)
}

func (o *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(o.PopLong()))
}

func (o *OperandStack) PushRef(val *heap.Object) {
	o.slots[o.size].ref = val
	o.valIdx[o.size] = true
	o.size++
}

func (o *OperandStack) PopRef() *heap.Object {
	o.size--
	o.valIdx[o.size] = false
	ref := o.slots[o.size].ref
	// 帮助Go的垃圾收集器回收Object结构体实例
	o.slots[o.size].ref = nil
	return ref
}

func (o *OperandStack) PushSlot(slot Slot) {
	o.slots[o.size] = slot
	o.valIdx[o.size] = true
	o.size++
}

func (o *OperandStack) PopSlot() Slot {
	o.size--
	o.valIdx[o.size] = false
	return o.slots[o.size]
}

func (o *OperandStack) PushTwoSlot(slot1, slot2 Slot) {
	o.slots[o.size] = slot1
	o.valIdx[o.size] = true
	o.size++
	o.slots[o.size] = slot2
	o.size++
}

func (o *OperandStack) IsTopSlotStartOfVal() bool {
	return o.valIdx[o.size-1]
}

func (o *OperandStack) GetRefFromTop(idx uint) *heap.Object {
	return o.slots[o.size-1-idx].ref
}
