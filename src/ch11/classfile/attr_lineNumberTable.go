package classfile

import "go-jvm/src/ch11/utils"

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

func (attr *LineNumberTableAttribute) readInfo(cr *utils.ByteCodeReader) {
	tableLen := cr.ReadUint16()
	attr.lineNumberTable = make([]*LineNumberTableEntry, tableLen)
	for i := range attr.lineNumberTable {
		attr.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    cr.ReadUint16(),
			lineNumber: cr.ReadUint16(),
		}
	}
}

func (attr *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(attr.lineNumberTable) - 1; i >= 0; i-- {
		if pc >= int(attr.lineNumberTable[i].startPc) {
			return int(attr.lineNumberTable[i].lineNumber)
		}
	}
	return -1
}
