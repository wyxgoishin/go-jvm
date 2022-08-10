package classfile

import "go-jvm/src/utils"

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/

type ConstantMethodType struct {
	descriptorIndex uint16
}

func (cpInfo *ConstantMethodType) readInfo(reader *utils.ByteCodeReader) {
	cpInfo.descriptorIndex = reader.ReadUint16()
}

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (cpInfo *ConstantMethodHandleInfo) readInfo(reader *utils.ByteCodeReader) {
	cpInfo.referenceKind = reader.ReadUint8()
	cpInfo.referenceIndex = reader.ReadUint16()
}

type ConstantInvokeDynamicInfo struct {
	boostrapMethodAttrIndex uint16
	nameAndType             uint16
}

func (cpInfo *ConstantInvokeDynamicInfo) readInfo(reader *utils.ByteCodeReader) {
	cpInfo.boostrapMethodAttrIndex = reader.ReadUint16()
	cpInfo.nameAndType = reader.ReadUint16()
}
