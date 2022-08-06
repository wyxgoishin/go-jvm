package runtimeDataArea

import (
	"math"
)

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack <= 0 {
		return nil
	}
	return &OperandStack{
		slots: make([]Slot, maxStack),
	}
}

func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].num = val
	o.size++
}

func (o *OperandStack) PopInt() int32 {
	o.size--
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

func (o *OperandStack) PushRef(val *Object) {
	o.slots[o.size].ref = val
	o.size++
}

func (o *OperandStack) PopRef() *Object {
	o.size--
	ref := o.slots[o.size].ref
	// 帮助Go的垃圾收集器回收Object结构体实例
	o.slots[o.size].ref = nil
	return ref
}
