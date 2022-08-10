package classfile

import "go-jvm/src/utils"

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (attr *ConstantValueAttribute) readInfo(reader *utils.ByteCodeReader) {
	attr.constantValueIndex = reader.ReadUint16()
}

// getter
func (attr *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return attr.constantValueIndex
}
