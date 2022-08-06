package classfile

import "go-jvm/src/ch05/utils"

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

func (c *ConstantStringInfo) readInfo(cr *utils.ByteCodeReader) {
	c.stringIndex = cr.ReadUint16()
}

func (c *ConstantStringInfo) String() string {
	return c.constantPool.getUtf8(c.stringIndex)
}
