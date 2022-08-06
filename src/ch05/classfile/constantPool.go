package classfile

import "go-jvm/src/ch05/utils"

type ConstantPool []ConstantInfo

func readConstantPool(cr *utils.ByteCodeReader) ConstantPool {
	count := int(cr.ReadUint16())
	constantPool := make([]ConstantInfo, count)
	// skip zero
	for i := 1; i < count; i++ {
		constantPool[i] = readConstantInfo(cr, constantPool)
		switch constantPool[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return constantPool
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if ci := cp[index]; ci != nil {
		return cp[index]
	}
	panic("Invalid constant pool index")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	ci := cp[index].(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(ci.nameIndex)
	_type := cp.getUtf8(ci.descriptorIndex)
	return name, _type
}

func (cp ConstantPool) getClassName(index uint16) string {
	ci := cp.getConstantInfo(index).(*ConstantClassInfo)
	return ci.Name()
}

func (cp ConstantPool) getUtf8(index uint16) string {
	ci := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return ci.str
}
