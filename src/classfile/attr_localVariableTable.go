package classfile

import "go-jvm/src/utils"

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (attr *LocalVariableTableAttribute) readInfo(reader *utils.ByteCodeReader) {
	tableLen := reader.ReadUint16()
	attr.localVariableTable = make([]*LocalVariableTableEntry, tableLen)
	for i := range attr.localVariableTable {
		attr.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.ReadUint16(),
			length:          reader.ReadUint16(),
			nameIndex:       reader.ReadUint16(),
			descriptorIndex: reader.ReadUint16(),
			index:           reader.ReadUint16(),
		}
	}
}
