package classfile

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

func (c *CodeAttribute) readInfo(cr *ClassReader) {
	c.maxStack = cr.readUint16()
	c.maxLocals = cr.readUint16()
	codeLength := cr.readUint32()
	c.code = cr.readBytes(codeLength)
	c.exceptionTable = readExceptionTable(cr)
	c.attributes = readAttributes(cr, c.constantPool)
}

func readExceptionTable(cr *ClassReader) []*ExceptionTableEntry {
	tableLen := cr.readUint16()
	table := make([]*ExceptionTableEntry, tableLen)
	for i := range table {
		table[i] = &ExceptionTableEntry{
			startPc:   cr.readUint16(),
			endPc:     cr.readUint16(),
			handlerPc: cr.readUint16(),
			catchType: cr.readUint16(),
		}
	}
	return table
}
