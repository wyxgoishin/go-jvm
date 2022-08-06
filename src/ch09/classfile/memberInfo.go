package classfile

import (
	"go-jvm/src/ch09/utils"
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

func readMemebers(cr *utils.ByteCodeReader, cp ConstantPool) []*MemberInfo {
	count := cr.ReadUint16()
	memberInfos := make([]*MemberInfo, count)
	for i := range memberInfos {
		memberInfos[i] = readMember(cr, cp)
	}
	return memberInfos
}

func readMember(cr *utils.ByteCodeReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    cp,
		accessFlags:     cr.ReadUint16(),
		nameIndex:       cr.ReadUint16(),
		descriptorIndex: cr.ReadUint16(),
		attributes:      readAttributes(cr, cp),
	}
}

func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.accessFlags
}

func (mi *MemberInfo) Name() string {
	return mi.constantPool.getUtf8(mi.nameIndex)
}

func (mi *MemberInfo) Descriptor() string {
	return mi.constantPool.getUtf8(mi.descriptorIndex)
	//return strings.Replace(mi.constantPool.getUtf8(mi.descriptorIndex), "/", ".", -1)
}

func (mi *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (mi *MemberInfo) ConstantValueAttributes() *ConstantValueAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
