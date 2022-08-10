package classfile

import (
	"go-jvm/src/utils"
)

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MemberInfo struct {
	constantPool    ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *utils.ByteCodeReader, cp ConstantPool) []*MemberInfo {
	count := reader.ReadUint16()
	memberInfos := make([]*MemberInfo, count)
	for i := range memberInfos {
		memberInfos[i] = readMember(reader, cp)
	}
	return memberInfos
}

func readMember(reader *utils.ByteCodeReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    cp,
		accessFlags:     reader.ReadUint16(),
		nameIndex:       reader.ReadUint16(),
		descriptorIndex: reader.ReadUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

// getters
func (cpInfo *MemberInfo) AccessFlags() uint16 {
	return cpInfo.accessFlags
}

func (cpInfo *MemberInfo) Name() string {
	return cpInfo.constantPool.GetUtf8(cpInfo.nameIndex)
}

func (cpInfo *MemberInfo) Descriptor() string {
	return cpInfo.constantPool.GetUtf8(cpInfo.descriptorIndex)
}

func (cpInfo *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range cpInfo.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (cpInfo *MemberInfo) ConstantValueAttributes() *ConstantValueAttribute {
	for _, attrInfo := range cpInfo.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (cpInfo *MemberInfo) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range cpInfo.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}
