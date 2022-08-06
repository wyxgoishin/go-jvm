package heap

import "go-jvm/src/ch06/classfile"

type Field struct {
	ClassMember
	slotId          uint
	constValueIndex uint
}

func newField(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = new(Field)
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (f *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttributes(); valAttr != nil {
		f.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (f *Field) isLongOrDouble() bool {
	return f.descriptor == "J" || f.descriptor == "D"
}

func (f *Field) SlotId() uint {
	return f.slotId
}

func (f *Field) ConstValueIndex() uint {
	return f.constValueIndex
}

func (f *Field) Class() *Class {
	return f.class
}

func (f *Field) AccessFlags() uint16 {
	return f.accessFlags
}

func (f *Field) Descriptor() string {
	return f.descriptor
}
