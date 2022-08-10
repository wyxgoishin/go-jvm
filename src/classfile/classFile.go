package classfile

import (
	"fmt"
	"go-jvm/src/utils"
)

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	magicNumber  uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (*ClassFile, error) {
	cr := utils.NewByteCodeReader(classData)
	cf := new(ClassFile)
	cf.read(cr)
	return cf, nil
}

func (clsFile *ClassFile) read(reader *utils.ByteCodeReader) {
	clsFile.readAndCheckMagic(reader)
	clsFile.readAndCheckVersion(reader)
	clsFile.constantPool = readConstantPool(reader)
	clsFile.accessFlags = reader.ReadUint16()
	clsFile.thisClass = reader.ReadUint16()
	clsFile.superClass = reader.ReadUint16()
	clsFile.interfaces = reader.ReadUint16s()
	clsFile.fields = readMembers(reader, clsFile.constantPool)
	clsFile.methods = readMembers(reader, clsFile.constantPool)
	clsFile.attributes = readAttributes(reader, clsFile.constantPool)
}

func (clsFile *ClassFile) readAndCheckMagic(reader *utils.ByteCodeReader) {
	clsFile.magicNumber = reader.ReadUint32()
	if clsFile.magicNumber != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

//ToDo: check version
func (clsFile *ClassFile) readAndCheckVersion(reader *utils.ByteCodeReader) {
	clsFile.minorVersion = reader.ReadUint16()
	clsFile.majorVersion = reader.ReadUint16()
	switch clsFile.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if clsFile.minorVersion == 0 {
			return
		}
	default:
		panic(fmt.Sprintf("java.lang.UnsupportedClassVersionError Version: %d.%d", clsFile.majorVersion, clsFile.minorVersion))
	}
}

// getters
func (clsFile *ClassFile) MinorVersion() uint16 {
	return clsFile.minorVersion
}

func (clsFile *ClassFile) MajorVersion() uint16 {
	return clsFile.majorVersion
}

func (clsFile *ClassFile) ConstantPool() ConstantPool {
	return clsFile.constantPool
}

func (clsFile *ClassFile) AccessFlags() uint16 {
	return clsFile.accessFlags
}

func (clsFile *ClassFile) Fields() []*MemberInfo {
	return clsFile.fields
}

func (clsFile *ClassFile) Methods() []*MemberInfo {
	return clsFile.methods
}

func (clsFile *ClassFile) ClassName() string {
	return clsFile.constantPool.getClassName(clsFile.thisClass)
}

func (clsFile *ClassFile) SuperClassName() string {
	if clsFile.superClass > 0 {
		return clsFile.constantPool.getClassName(clsFile.superClass)
	}
	return ""
}

func (clsFile *ClassFile) InterfaceNames() []string {
	var interfaceNames []string
	for _, inter := range clsFile.interfaces {
		interfaceNames = append(interfaceNames, clsFile.constantPool.getClassName(inter))
	}
	return interfaceNames
}

func (clsFile *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range clsFile.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}
