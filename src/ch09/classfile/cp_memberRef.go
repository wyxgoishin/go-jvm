package classfile

import "go-jvm/src/ch09/utils"

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantFieldRefInfo struct {
	ConstantMemberInfo
}

/*
CONSTANT_Methodref_info {
u1 tag;
u2 class_index;
u2 name_and_type_index;
}
*/
type ConstantMethodRefInfo struct {
	ConstantMemberInfo
}

/*
CONSTANT_InterfaceMethodref_info {
u1 tag;
u2 class_index;
u2 name_and_type_index;
}
*/
type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberInfo
}

type ConstantMemberInfo struct {
	constantPool     ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (c *ConstantMemberInfo) readInfo(cr *utils.ByteCodeReader) {
	c.classIndex = cr.ReadUint16()
	c.nameAndTypeIndex = cr.ReadUint16()
}

func (c *ConstantMemberInfo) ClassName() string {
	return c.constantPool.getClassName(c.classIndex)
}

func (c *ConstantMemberInfo) NameAndDescriptor() (string, string) {
	return c.constantPool.getNameAndType(c.nameAndTypeIndex)
}
