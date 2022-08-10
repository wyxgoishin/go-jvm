package classfile

import "go-jvm/src/utils"

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
	u1 tag;
	u2 class_index;
	u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
	u1 tag;
	u2 class_index;
	u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
	u1 tag;
	u2 class_index;
	u2 name_and_type_index;
}
*/

type ConstantFieldRefInfo struct {
	ConstantMemberInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberInfo
}

type ConstantMemberInfo struct {
	constantPool     ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (cpInfo *ConstantMemberInfo) readInfo(reader *utils.ByteCodeReader) {
	cpInfo.classIndex = reader.ReadUint16()
	cpInfo.nameAndTypeIndex = reader.ReadUint16()
}

func (cpInfo *ConstantMemberInfo) ClassName() string {
	return cpInfo.constantPool.getClassName(cpInfo.classIndex)
}

func (cpInfo *ConstantMemberInfo) NameAndDescriptor() (string, string) {
	return cpInfo.constantPool.getNameAndType(cpInfo.nameAndTypeIndex)
}
