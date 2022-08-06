package classfile

import "go-jvm/src/ch09/utils"

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (m *MarkerAttribute) readInfo(cr *utils.ByteCodeReader) {
	// 二者都是标记信息，里面没有内容
}
