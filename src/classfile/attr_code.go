package classfile

import (
	"go-jvm/src/utils"
)

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	constantPool   ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

// readers
func (attr *CodeAttribute) readInfo(reader *utils.ByteCodeReader) {
	attr.maxStack = reader.ReadUint16()
	attr.maxLocals = reader.ReadUint16()
	codeLength := reader.ReadUint32()
	attr.code = reader.ReadBytes(codeLength)
	attr.exceptionTable = readExceptionTable(reader)
	attr.attributes = readAttributes(reader, attr.constantPool)
}

func readExceptionTable(reader *utils.ByteCodeReader) []*ExceptionTableEntry {
	tableLen := reader.ReadUint16()
	table := make([]*ExceptionTableEntry, tableLen)
	for i := range table {
		table[i] = &ExceptionTableEntry{
			startPc:   reader.ReadUint16(),
			endPc:     reader.ReadUint16(),
			handlerPc: reader.ReadUint16(),
			catchType: reader.ReadUint16(),
		}
	}
	return table
}

// getters of CodeAttribute
func (attr *CodeAttribute) MaxLocals() uint16 {
	return attr.maxLocals
}

func (attr *CodeAttribute) MaxStack() uint16 {
	return attr.maxStack
}

func (attr *CodeAttribute) Code() []byte {
	return attr.code
}

func (attr *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return attr.exceptionTable
}

// getters of ExceptionTableEntry
func (entry *ExceptionTableEntry) StartPc() uint16 {
	return entry.startPc
}

func (entry *ExceptionTableEntry) EndPc() uint16 {
	return entry.endPc
}

func (entry *ExceptionTableEntry) HandlerPc() uint16 {
	return entry.handlerPc
}

func (entry *ExceptionTableEntry) CatchType() uint16 {
	return entry.catchType
}
