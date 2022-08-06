package classfile

import (
	"fmt"
	"go-jvm/src/ch08/utils"
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

func (cf *ClassFile) read(cr *utils.ByteCodeReader) {
	cf.readAndCheckMagic(cr)
	cf.readAndCheckVersion(cr)
	cf.constantPool = readConstantPool(cr)
	cf.accessFlags = cr.ReadUint16()
	cf.thisClass = cr.ReadUint16()
	cf.superClass = cr.ReadUint16()
	cf.interfaces = cr.ReadUint16s()
	cf.fields = readMemebers(cr, cf.constantPool)
	cf.methods = readMemebers(cr, cf.constantPool)
	cf.attributes = readAttributes(cr, cf.constantPool)
}

func (cf *ClassFile) readAndCheckMagic(cr *utils.ByteCodeReader) {
	cf.magicNumber = cr.ReadUint32()
	if cf.magicNumber != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(cr *utils.ByteCodeReader) {
	cf.minorVersion = cr.ReadUint16()
	cf.majorVersion = cr.ReadUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	default:
		panic(fmt.Sprintf("java.lang.UnsupportedClassVersionError Version: %d.%d", cf.majorVersion, cf.minorVersion))
	}
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""
}

func (cf *ClassFile) InterfaceNames() []string {
	var interfaceNames []string
	for _, inter := range cf.interfaces {
		interfaceNames = append(interfaceNames, cf.constantPool.getClassName(inter))
	}
	return interfaceNames
}
