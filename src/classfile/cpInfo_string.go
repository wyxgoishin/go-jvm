package classfile

import "go-jvm/src/utils"

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	constantPool ConstantPool
	stringIndex  uint16
}

func (cpInfo *ConstantStringInfo) readInfo(reader *utils.ByteCodeReader) {
	cpInfo.stringIndex = reader.ReadUint16()
}

func (cpInfo *ConstantStringInfo) String() string {
	return cpInfo.constantPool.GetUtf8(cpInfo.stringIndex)
}
