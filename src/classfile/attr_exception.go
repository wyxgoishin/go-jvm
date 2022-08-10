package classfile

import "go-jvm/src/utils"

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (attr *ExceptionAttribute) readInfo(reader *utils.ByteCodeReader) {
	attr.exceptionIndexTable = reader.ReadUint16s()
}
