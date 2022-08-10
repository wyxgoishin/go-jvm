package heap

import "fmt"

type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (obj *Object) Fields() Slots {
	return obj.data.(Slots)
}

func (obj *Object) IsInstanceOf(dst *Class) bool {
	return dst.isAssignableFrom(obj.class)
}

func (obj *Object) Class() *Class {
	return obj.class
}

func (obj *Object) SetInstanceRefVar(name, descriptor string, ref *Object) {
	field := obj.Class().getField(name, descriptor, false)
	slots := obj.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (obj *Object) GetInstanceRefVar(name, descriptor string) *Object {
	if obj == nil {
		return nil
	}
	filed := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	return slots.GetRef(filed.slotId)
}

func (obj *Object) Extra() interface{} {
	return obj.extra
}

func (obj *Object) SetExtra(extra interface{}) {
	obj.extra = extra
}

func (obj *Object) Int() int32 {
	return obj.data.(int32)
}

func (obj *Object) Long() int64 {
	return obj.data.(int64)
}

func (obj *Object) Float() float32 {
	return obj.data.(float32)
}

func (obj *Object) Double() float64 {
	return obj.data.(float64)
}

func (obj *Object) Clone() *Object {
	return &Object{
		class: obj.class,
		data:  obj.cloneData(), // 注意 data 可能是切片，所以不能直接赋值
	}
}

func (obj *Object) Data() interface{} {
	return obj.data
}

func (obj *Object) Print(onlyRef bool) {
	if obj != nil && obj.data != nil {
		switch obj.data.(type) {
		case []int8, []int16, []uint16, []int32, []int64, []float32, []float64, int8, int16, uint16, int32, int64, float32, float64:
			fmt.Printf("%v\n", obj.data)
		case Slots:
			slots := obj.data.(Slots)
			for _, slot := range slots {
				slot.Print(onlyRef)
			}
		case []*Object:
			objs := obj.data.([]*Object)
			for _, subObj := range objs {
				subObj.Print(false)
			}
		default:
			// ToDo
			panic("Implementing me")
		}
	}
}

func (obj *Object) cloneData() interface{} {
	switch obj.data.(type) {
	case []int8:
		src := obj.data.([]int8)
		dst := make([]int8, len(src))
		copy(dst, src)
		return dst
	case []int16:
		src := obj.data.([]int16)
		dst := make([]int16, len(src))
		copy(dst, src)
		return dst
	case []uint16:
		src := obj.data.([]uint16)
		dst := make([]uint16, len(src))
		copy(dst, src)
		return dst
	case []int32:
		src := obj.data.([]int32)
		dst := make([]int32, len(src))
		copy(dst, src)
		return dst
	case []int64:
		src := obj.data.([]int64)
		dst := make([]int64, len(src))
		copy(dst, src)
		return dst
	case []float32:
		src := obj.data.([]float32)
		dst := make([]float32, len(src))
		copy(dst, src)
		return dst
	case []float64:
		src := obj.data.([]float64)
		dst := make([]float64, len(src))
		copy(dst, src)
		return dst
	case []*Object:
		src := obj.data.([]*Object)
		dst := make([]*Object, len(src))
		copy(dst, src)
		return dst
	case []Slot:
		src := obj.data.([]*Slot)
		dst := make([]*Slot, len(src))
		copy(dst, src)
		return dst
	default:
		return obj.data
	}
}
