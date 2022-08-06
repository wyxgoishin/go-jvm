package classfile

import "go-jvm/src/ch06/utils"

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (l *LineNumberTableAttribute) readInfo(cr *utils.ByteCodeReader) {
	tableLen := cr.ReadUint16()
	l.lineNumberTable = make([]*LineNumberTableEntry, tableLen)
	for i := range l.lineNumberTable {
		l.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    cr.ReadUint16(),
			lineNumber: cr.ReadUint16(),
		}
	}
}
