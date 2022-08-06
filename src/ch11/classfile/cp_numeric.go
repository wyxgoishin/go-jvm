package classfile

import (
	"go-jvm/src/ch11/utils"
	"math"
)

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (c *ConstantIntegerInfo) readInfo(cr *utils.ByteCodeReader) {
	bytes := cr.ReadUint32()
	c.val = int32(bytes)
}

func (c *ConstantIntegerInfo) Value() int32 {
	return c.val
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

func (c *ConstantFloatInfo) readInfo(cr *utils.ByteCodeReader) {
	bytes := cr.ReadUint32()
	c.val = float32(bytes)
}

func (c *ConstantFloatInfo) Value() float32 {
	return c.val
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

func (c *ConstantLongInfo) readInfo(cr *utils.ByteCodeReader) {
	bytes := cr.ReadUint64()
	c.val = int64(bytes)
}

func (c *ConstantLongInfo) Value() int64 {
	return c.val
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

func (c *ConstantDoubleInfo) readInfo(cr *utils.ByteCodeReader) {
	bytes := cr.ReadUint64()
	c.val = math.Float64frombits(bytes)
}

func (c *ConstantDoubleInfo) Value() float64 {
	return c.val
}
