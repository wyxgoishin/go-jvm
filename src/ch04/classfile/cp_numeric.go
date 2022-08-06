package classfile

import "math"

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (c *ConstantIntegerInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint32()
	c.val = int32(bytes)
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (c *ConstantFloatInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint32()
	c.val = float32(bytes)
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	val int64
}

func (c *ConstantLongInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint64()
	c.val = int64(bytes)
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (c *ConstantDoubleInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint64()
	c.val = math.Float64frombits(bytes)
}
