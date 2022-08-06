package classfile

import "go-jvm/src/ch11/utils"

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

func (attr *CodeAttribute) readInfo(cr *utils.ByteCodeReader) {
	attr.maxStack = cr.ReadUint16()
	attr.maxLocals = cr.ReadUint16()
	codeLength := cr.ReadUint32()
	attr.code = cr.ReadBytes(codeLength)
	attr.exceptionTable = readExceptionTable(cr)
	attr.attributes = readAttributes(cr, attr.constantPool)
}

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

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

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

func readExceptionTable(cr *utils.ByteCodeReader) []*ExceptionTableEntry {
	tableLen := cr.ReadUint16()
	table := make([]*ExceptionTableEntry, tableLen)
	for i := range table {
		table[i] = &ExceptionTableEntry{
			startPc:   cr.ReadUint16(),
			endPc:     cr.ReadUint16(),
			handlerPc: cr.ReadUint16(),
			catchType: cr.ReadUint16(),
		}
	}
	return table
}
