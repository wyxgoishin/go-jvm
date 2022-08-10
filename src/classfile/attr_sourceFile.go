package classfile

import (
	"go-jvm/src/utils"
)

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
	constantPool    ConstantPool
	sourceFileIndex uint16
}

func (attr *SourceFileAttribute) readInfo(reader *utils.ByteCodeReader) {
	attr.sourceFileIndex = reader.ReadUint16()
}

func (attr *SourceFileAttribute) FileName() string {
	return attr.constantPool.GetUtf8(attr.sourceFileIndex)
}
