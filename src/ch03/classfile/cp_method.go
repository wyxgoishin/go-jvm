package classfile

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodType struct {
	descriptorIndex uint16
}

func (c *ConstantMethodType) readInfo(cr *ClassReader) {
	c.descriptorIndex = cr.readUint16()
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

func (c *ConstantMethodHandleInfo) readInfo(cr *ClassReader) {
	c.referenceKind = cr.readUint8()
	c.referenceIndex = cr.readUint16()
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

func (c *ConstantInvokeDynamicInfo) readInfo(cr *ClassReader) {
	c.boostrapMethodAttrIndex = cr.readUint16()
	c.nameAndType = cr.readUint16()
}
