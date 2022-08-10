package classfile

import (
	"go-jvm/src/utils"
)

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
	constantPool ConstantPool
	nameIndex    uint16
}

func (cpInfo *ConstantClassInfo) readInfo(reader *utils.ByteCodeReader) {
	cpInfo.nameIndex = reader.ReadUint16()
}

func (cpInfo *ConstantClassInfo) Name() string {
	return cpInfo.constantPool.GetUtf8(cpInfo.nameIndex)
}
