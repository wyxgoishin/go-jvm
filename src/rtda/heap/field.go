package heap

import (
	"go-jvm/src/classfile"
)

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

func (field *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttributes(); valAttr != nil {
		field.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (field *Field) isLongOrDouble() bool {
	return field.descriptor == "J" || field.descriptor == "D"
}

func (field *Field) SlotId() uint {
	return field.slotId
}

func (field *Field) ConstValueIndex() uint {
	return field.constValueIndex
}

func (field *Field) Class() *Class {
	return field.class
}

func (field *Field) IsVolatile() bool {
	return field.accessFlags&ACC_VOLATILE > 0
}

func (field *Field) IsTransient() bool {
	return field.accessFlags&ACC_TRANSIENT > 0
}

func (field *Field) IsEnum() bool {
	return field.accessFlags&ACC_ENUM > 0
}
