package classfile

import "go-jvm/src/ch09/utils"

const (
	CONSTANT_Class              = 7
	CONSTANT_FieldRef           = 9
	CONSTANT_MethodRef          = 10
	CONSTANT_InterfaceMethodRef = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

/*
cp_info {
    u1 tag;
    u1 info[];
}
*/
type ConstantInfo interface {
	readInfo(cr *utils.ByteCodeReader)
}

func readConstantInfo(cr *utils.ByteCodeReader, cp ConstantPool) ConstantInfo {
	tag := cr.ReadUint8()
	ci := newConstantInfo(tag, cp)
	ci.readInfo(cr)
	return ci
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Class:
		return &ConstantClassInfo{constantPool: cp}
	case CONSTANT_FieldRef:
		return &ConstantFieldRefInfo{ConstantMemberInfo{constantPool: cp}}
	case CONSTANT_MethodRef:
		return &ConstantMethodRefInfo{ConstantMemberInfo{constantPool: cp}}
	case CONSTANT_InterfaceMethodRef:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberInfo{constantPool: cp}}
	case CONSTANT_String:
		return &ConstantStringInfo{constantPool: cp}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodType{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError:  constant pool tag")
	}
}
