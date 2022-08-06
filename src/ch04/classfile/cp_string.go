package classfile

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

func (c *ConstantStringInfo) readInfo(cr *ClassReader) {
	c.stringIndex = cr.readUint16()
}

func (c *ConstantStringInfo) String() string {
	return c.constantPool.getUtf8(c.stringIndex)
}
