package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(cr *ClassReader)
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

func readAttributes(cr *ClassReader, cp ConstantPool) []AttributeInfo {
	count := cr.readUint16()
	attributeInfos := make([]AttributeInfo, count)
	for i := range attributeInfos {
		attributeInfos[i] = readAttribute(cr, cp)
	}
	return nil
}

func readAttribute(cr *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIdx := cr.readUint16()
	attrName := cp.getUtf8(attrNameIdx)
	attrLen := cr.readUint32()
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
