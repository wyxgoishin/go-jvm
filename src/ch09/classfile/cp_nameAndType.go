package classfile

import "go-jvm/src/ch09/utils"

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

func (c *ConstantNameAndTypeInfo) readInfo(cr *utils.ByteCodeReader) {
	c.nameIndex = cr.ReadUint16()
	c.descriptorIndex = cr.ReadUint16()
}
