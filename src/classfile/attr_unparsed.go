package classfile

import "go-jvm/src/utils"

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (attr *UnparsedAttribute) readInfo(reader *utils.ByteCodeReader) {
	attr.info = reader.ReadBytes(attr.length)
}
