package classfile

import "go-jvm/src/utils"

//ToDo: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	*MarkerAttribute
}

type SyntheticAttribute struct {
	*MarkerAttribute
}

type MarkerAttribute struct{}

func (attr *MarkerAttribute) readInfo(reader *utils.ByteCodeReader) {
	// 二者都是标记信息，里面没有内容
}
