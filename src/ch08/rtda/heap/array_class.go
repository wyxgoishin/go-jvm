package heap

func (c *Class) NewArray(count uint) *Object {
	if !c.IsArray() {
		panic("Not array clas: " + c.name)
	}
	switch c.Name() {
	case "[Z":
		return &Object{c, make([]int8, count)} // bool
	case "[B":
		return &Object{c, make([]int8, count)}
	case "[S":
		return &Object{c, make([]uint16, count)}
	case "[C":
		return &Object{c, make([]int16, count)}
	case "[I":
		return &Object{c, make([]int32, count)}
	case "[F":
		return &Object{c, make([]float32, count)}
	case "[J":
		return &Object{c, make([]int64, count)}
	case "[D":
		return &Object{c, make([]float64, count)}
	default:
		return &Object{c, make([]*Object, count)}
	}
}
