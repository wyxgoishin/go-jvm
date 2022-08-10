package classfile

import "go-jvm/src/utils"

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

func (attr *LineNumberTableAttribute) readInfo(reader *utils.ByteCodeReader) {
	tableLen := reader.ReadUint16()
	attr.lineNumberTable = make([]*LineNumberTableEntry, tableLen)
	for i := range attr.lineNumberTable {
		attr.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.ReadUint16(),
			lineNumber: reader.ReadUint16(),
		}
	}
}

func (attr *LineNumberTableAttribute) GetLineNumber(pc uint) int {
	for i := len(attr.lineNumberTable) - 1; i >= 0; i-- {
		if pc >= uint(attr.lineNumberTable[i].startPc) {
			return int(attr.lineNumberTable[i].lineNumber)
		}
	}
	return -1
}
