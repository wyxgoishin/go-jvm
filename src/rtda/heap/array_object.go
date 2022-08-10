package heap

import "fmt"

func (obj *Object) Bytes() []int8 {
	return obj.data.([]int8)
}

func (obj *Object) Shorts() []int16 {
	return obj.data.([]int16)
}

func (obj *Object) Chars() []uint16 {
	return obj.data.([]uint16)
}

func (obj *Object) Ints() []int32 {
	return obj.data.([]int32)
}

func (obj *Object) Floats() []float32 {
	return obj.data.([]float32)
}

func (obj *Object) Longs() []int64 {
	return obj.data.([]int64)
}

func (obj *Object) Doubles() []float64 {
	return obj.data.([]float64)
}

func (obj *Object) Slots() Slots {
	return obj.data.(Slots)
}

func (obj *Object) Refs() []*Object {
	return obj.data.([]*Object)
}

func (obj *Object) ArrayLength() int {
	switch obj.data.(type) {
	case []int8:
		return len(obj.Bytes())
	case []int16:
		return len(obj.Shorts())
	case []uint16:
		return len(obj.Chars())
	case []int32:
		return len(obj.Ints())
	case []float32:
		return len(obj.Floats())
	case []int64:
		return len(obj.Longs())
	case []float64:
		return len(obj.Doubles())
	case []*Object:
		return len(obj.Refs())
	default:
		panic(fmt.Sprintf("Unexpected array elem type: %T", obj.data))
	}
}

func ArrayCopy(src *Object, srcPos int32, dst *Object, dstPos, length int32) {
	switch src.data.(type) {
	case []int8:
		_src := src.data.([]int8)[srcPos : srcPos+length]
		_dst := dst.data.([]int8)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int16:
		_src := src.data.([]int16)[srcPos : srcPos+length]
		_dst := dst.data.([]int16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []uint16:
		_src := src.data.([]uint16)[srcPos : srcPos+length]
		_dst := dst.data.([]uint16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int32:
		_src := src.data.([]int32)[srcPos : srcPos+length]
		_dst := dst.data.([]int32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float32:
		_src := src.data.([]float32)[srcPos : srcPos+length]
		_dst := dst.data.([]float32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float64:
		_src := src.data.([]float64)[srcPos : srcPos+length]
		_dst := dst.data.([]float64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []*Object:
		_src := src.data.([]*Object)[srcPos : srcPos+length]
		_dst := dst.data.([]*Object)[dstPos : dstPos+length]
		copy(_dst, _src)
	}
}
