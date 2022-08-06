package classfile

import (
	"go-jvm/src/ch10/utils"
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

func (c *ConstantClassInfo) readInfo(cr *utils.ByteCodeReader) {
	c.nameIndex = cr.ReadUint16()
}

func (c *ConstantClassInfo) Name() string {
	return c.constantPool.getUtf8(c.nameIndex)
}
