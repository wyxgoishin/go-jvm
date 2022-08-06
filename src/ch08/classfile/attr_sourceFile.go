package classfile

import "go-jvm/src/ch08/utils"

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

func (s *SourceFileAttribute) readInfo(cr *utils.ByteCodeReader) {
	s.sourceFileIndex = cr.ReadUint16()
}

func (s *SourceFileAttribute) FileName() string {
	return s.constantPool.getUtf8(s.sourceFileIndex)
}
