package classfile

import "strings"

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

func readMemebers(cr *ClassReader, cp ConstantPool) []*MemberInfo {
	count := cr.readUint16()
	memberInfos := make([]*MemberInfo, count)
	for i := range memberInfos {
		memberInfos[i] = readMember(cr, cp)
	}
	return memberInfos
}

func readMember(cr *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    cp,
		accessFlags:     cr.readUint16(),
		nameIndex:       cr.readUint16(),
		descriptorIndex: cr.readUint16(),
		attributes:      readAttributes(cr, cp),
	}
}

func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.accessFlags
}

func (mi *MemberInfo) Name() string {
	return mi.constantPool.getUtf8(mi.nameIndex)
}

func (mi *MemberInfo) Description() string {
	return strings.Replace(mi.constantPool.getUtf8(mi.descriptorIndex), "/", ".", -1)
}
