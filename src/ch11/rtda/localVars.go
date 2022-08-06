package rtda

import (
	"go-jvm/src/ch11/rtda/heap"
	"math"
)

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals <= 0 {
		return nil
	}
	return make([]Slot, maxLocals)
}

func (localVars LocalVars) SetInt(index uint, val int32) {
	localVars[index].num = val
}

func (localVars LocalVars) GetInt(index uint) int32 {
	return localVars[index].num
}

func (localVars LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	localVars[index].num = int32(bits)
}

func (localVars LocalVars) GetFloat(index uint) float32 {
	bits := uint32(localVars[index].num)
	return math.Float32frombits(bits)
}

func (localVars LocalVars) SetLong(index uint, val int64) {
	localVars[index].num = int32(val)
	localVars[index+1].num = int32(val >> 32)
}

func (localVars LocalVars) GetLong(index uint) int64 {
	lower := int64(localVars[index].num)
	higher := int64(localVars[index+1].num) << 32
	return higher + lower
}

func (localVars LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	localVars.SetLong(index, int64(bits))
}

func (localVars LocalVars) GetDouble(index uint) float64 {
	bits := uint64(localVars.GetLong(index))
	return math.Float64frombits(bits)
}

func (localVars LocalVars) SetRef(index uint, ref *heap.Object) {
	localVars[index].ref = ref
}

func (localVars LocalVars) GetRef(index uint) *heap.Object {
	return localVars[index].ref
}

func (localVars LocalVars) SetSlot(index uint, slot Slot) {
	localVars[index] = slot
}

func (localVars *LocalVars) GetThis() *heap.Object {
	return localVars.GetRef(0)
}
