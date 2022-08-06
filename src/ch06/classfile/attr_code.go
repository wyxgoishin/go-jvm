package classfile

import "go-jvm/src/ch06/utils"

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

func (c *CodeAttribute) readInfo(cr *utils.ByteCodeReader) {
	c.maxStack = cr.ReadUint16()
	c.maxLocals = cr.ReadUint16()
	codeLength := cr.ReadUint32()
	c.code = cr.ReadBytes(codeLength)
	c.exceptionTable = readExceptionTable(cr)
	c.attributes = readAttributes(cr, c.constantPool)
}

func (c *CodeAttribute) MaxLocals() uint16 {
	return c.maxLocals
}

func (c *CodeAttribute) MaxStack() uint16 {
	return c.maxStack
}

func (c *CodeAttribute) Code() []byte {
	return c.code
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
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
