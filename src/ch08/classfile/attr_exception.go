package classfile

import "go-jvm/src/ch08/utils"

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

func (e *ExceptionAttribute) readInfo(cr *utils.ByteCodeReader) {
	e.exceptionIndexTable = cr.ReadUint16s()
}
