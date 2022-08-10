package classfile

import (
	"go-jvm/src/utils"
)

const (
	Code               = "Code"
	ConstantValue      = "ConstantValue"
	Deprecated         = "Deprecated"
	Exceptions         = "Exceptions"
	LineNumberTable    = "LineNumberTable"
	LocalVariableTable = "LocalVariableTable"
	SourceFile         = "SourceFile"
	Synthetic          = "Synthetic"
)

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *utils.ByteCodeReader)
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case Code:
		return &CodeAttribute{constantPool: cp}
	case ConstantValue:
		return &ConstantValueAttribute{}
	case Deprecated:
		return &DeprecatedAttribute{}
	case Exceptions:
		return &ExceptionAttribute{}
	case LineNumberTable:
		return &LineNumberTableAttribute{}
	case LocalVariableTable:
		return &LocalVariableTableAttribute{}
	case SourceFile:
		return &SourceFileAttribute{constantPool: cp}
	case Synthetic:
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}

func readAttributes(reader *utils.ByteCodeReader, cp ConstantPool) []AttributeInfo {
	count := reader.ReadUint16()
	attrInfos := make([]AttributeInfo, count)
	for i := range attrInfos {
		attrInfos[i] = readAttribute(reader, cp)
	}
	return attrInfos
}

func readAttribute(reader *utils.ByteCodeReader, cp ConstantPool) AttributeInfo {
	attrNameIdx := reader.ReadUint16()
	attrName := cp.GetUtf8(attrNameIdx)
	attrLen := reader.ReadUint32()
	attributeInfo := newAttributeInfo(attrName, attrLen, cp)
	attributeInfo.readInfo(reader)
	return attributeInfo
}
