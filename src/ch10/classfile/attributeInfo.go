package classfile

import "go-jvm/src/ch10/utils"

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(cr *utils.ByteCodeReader)
}

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

func readAttributes(cr *utils.ByteCodeReader, cp ConstantPool) []AttributeInfo {
	count := cr.ReadUint16()
	attributeInfos := make([]AttributeInfo, count)
	for i := range attributeInfos {
		attributeInfos[i] = readAttribute(cr, cp)
	}
	return attributeInfos
}

func readAttribute(cr *utils.ByteCodeReader, cp ConstantPool) AttributeInfo {
	attrNameIdx := cr.ReadUint16()
	attrName := cp.getUtf8(attrNameIdx)
	attrLen := cr.ReadUint32()
	attributeInfo := newAttributeInfo(attrName, attrLen, cp)
	attributeInfo.readInfo(cr)
	return attributeInfo
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
