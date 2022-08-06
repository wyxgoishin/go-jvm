package classfile

import "go-jvm/src/ch05/utils"

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

func (l *LocalVariableTableAttribute) readInfo(cr *utils.ByteCodeReader) {
	tableLen := cr.ReadUint16()
	l.localVariableTable = make([]*LocalVariableTableEntry, tableLen)
	for i := range l.localVariableTable {
		l.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         cr.ReadUint16(),
			length:          cr.ReadUint16(),
			nameIndex:       cr.ReadUint16(),
			descriptorIndex: cr.ReadUint16(),
			index:           cr.ReadUint16(),
		}
	}
}
