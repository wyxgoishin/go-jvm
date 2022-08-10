package classfile

import "go-jvm/src/utils"

type ConstantPool []ConstantInfo

func readConstantPool(reader *utils.ByteCodeReader) ConstantPool {
	count := int(reader.ReadUint16())
	constantPool := make([]ConstantInfo, count)
	// skip zero
	for i := 1; i < count; i++ {
		constantPool[i] = readConstantInfo(reader, constantPool)
		switch constantPool[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return constantPool
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cp[index]
	}
	panic("Invalid constant pool index")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	cpInfo := cp[index].(*ConstantNameAndTypeInfo)
	name := cp.GetUtf8(cpInfo.nameIndex)
	_type := cp.GetUtf8(cpInfo.descriptorIndex)
	return name, _type
}

func (cp ConstantPool) getClassName(index uint16) string {
	cpInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cpInfo.Name()
}

func (cp ConstantPool) GetUtf8(index uint16) string {
	cpInfo := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return cpInfo.str
}
