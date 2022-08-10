package classfile

import (
	"go-jvm/src/utils"
	"math"
)

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

type ConstantIntegerInfo struct {
	val int32
}

func (cpInfo *ConstantIntegerInfo) readInfo(reader *utils.ByteCodeReader) {
	bytes := reader.ReadUint32()
	cpInfo.val = int32(bytes)
}

func (cpInfo *ConstantIntegerInfo) Value() int32 {
	return cpInfo.val
}

type ConstantFloatInfo struct {
	val float32
}

func (cpInfo *ConstantFloatInfo) readInfo(reader *utils.ByteCodeReader) {
	bytes := reader.ReadUint32()
	cpInfo.val = float32(bytes)
}

func (cpInfo *ConstantFloatInfo) Value() float32 {
	return cpInfo.val
}

type ConstantLongInfo struct {
	val int64
}

func (cpInfo *ConstantLongInfo) readInfo(reader *utils.ByteCodeReader) {
	bytes := reader.ReadUint64()
	cpInfo.val = int64(bytes)
}

func (cpInfo *ConstantLongInfo) Value() int64 {
	return cpInfo.val
}

type ConstantDoubleInfo struct {
	val float64
}

func (cpInfo *ConstantDoubleInfo) readInfo(reader *utils.ByteCodeReader) {
	bytes := reader.ReadUint64()
	cpInfo.val = math.Float64frombits(bytes)
}

func (cpInfo *ConstantDoubleInfo) Value() float64 {
	return cpInfo.val
}
