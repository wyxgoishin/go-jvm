package classfile

import "go-jvm/src/ch05/utils"

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodType struct {
	descriptorIndex uint16
}

func (c *ConstantMethodType) readInfo(cr *utils.ByteCodeReader) {
	c.descriptorIndex = cr.ReadUint16()
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (c *ConstantMethodHandleInfo) readInfo(cr *utils.ByteCodeReader) {
	c.referenceKind = cr.ReadUint8()
	c.referenceIndex = cr.ReadUint16()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	boostrapMethodAttrIndex uint16
	nameAndType             uint16
}

func (c *ConstantInvokeDynamicInfo) readInfo(cr *utils.ByteCodeReader) {
	c.boostrapMethodAttrIndex = cr.ReadUint16()
	c.nameAndType = cr.ReadUint16()
}
