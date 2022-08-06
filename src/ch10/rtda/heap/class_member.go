package heap

import "go-jvm/src/ch10/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (clsMember *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	clsMember.accessFlags = memberInfo.AccessFlags()
	clsMember.name = memberInfo.Name()
	clsMember.descriptor = memberInfo.Descriptor()
}

func (clsMember *ClassMember) isAccessibleTo(other *Class) bool {
	if clsMember.IsPublic() {
		return true
	}
	this := clsMember.class
	if clsMember.IsProtected() {
		return this == other || other.IsSubClassOf(this) || this.GetPackageName() == other.GetPackageName()
	}
	if !clsMember.IsPrivate() {
		return this.GetPackageName() == other.GetPackageName()
	}
	return this == other
}

func (clsMember *ClassMember) AccessFlags() uint16 {
	return clsMember.accessFlags
}

func (clsMember *ClassMember) Name() string {
	return clsMember.name
}

func (clsMember *ClassMember) Descriptor() string {
	return clsMember.descriptor
}

func (clsMember *ClassMember) IsPublic() bool {
	return clsMember.accessFlags&ACC_PUBLIC > 0
}

func (clsMember *ClassMember) IsPrivate() bool {
	return clsMember.accessFlags&ACC_PRIVATE > 0
}

func (clsMember *ClassMember) IsProtected() bool {
	return clsMember.accessFlags&ACC_PROTECTED > 0
}

func (clsMember *ClassMember) IsStatic() bool {
	return clsMember.accessFlags&ACC_STATIC > 0
}

func (clsMember *ClassMember) IsFinal() bool {
	return clsMember.accessFlags&ACC_FINAL > 0
}

func (clsMember *ClassMember) IsSynthetic() bool {
	return clsMember.accessFlags&ACC_SYNTHETIC > 0
}
