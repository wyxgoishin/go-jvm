package classfile

import "go-jvm/src/ch09/utils"

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

func (u *UnparsedAttribute) readInfo(cr *utils.ByteCodeReader) {
	u.info = cr.ReadBytes(u.length)
}
