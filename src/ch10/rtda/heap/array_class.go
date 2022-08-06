package heap

func (cls *Class) NewArray(count uint) *Object {
	if !cls.IsArray() {
		panic("Not array clas: " + cls.name)
	}
	switch cls.Name() {
	case "[Z":
		return &Object{class: cls, data: make([]int8, count)} // bool
	case "[B":
		return &Object{class: cls, data: make([]int8, count)}
	case "[S":
		return &Object{class: cls, data: make([]int16, count)}
	case "[C":
		return &Object{class: cls, data: make([]uint16, count)}
	case "[I":
		return &Object{class: cls, data: make([]int32, count)}
	case "[F":
		return &Object{class: cls, data: make([]float32, count)}
	case "[J":
		return &Object{class: cls, data: make([]int64, count)}
	case "[D":
		return &Object{class: cls, data: make([]float64, count)}
	default:
		return &Object{class: cls, data: make([]*Object, count)}
	}
}
