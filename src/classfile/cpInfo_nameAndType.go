package classfile

import "go-jvm/src/utils"

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (cpInfo *ConstantNameAndTypeInfo) readInfo(reader *utils.ByteCodeReader) {
	cpInfo.nameIndex = reader.ReadUint16()
	cpInfo.descriptorIndex = reader.ReadUint16()
}
