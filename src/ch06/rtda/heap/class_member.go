package heap

import "go-jvm/src/ch06/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (c *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	c.accessFlags = memberInfo.AccessFlags()
	c.name = memberInfo.Name()
	c.descriptor = memberInfo.Descriptor()
}

func (c *ClassMember) isAccessibleTo(other *Class) bool {
	if IsPublic(c.accessFlags) {
		return true
	}
	this := c.class
	if IsProtected(this.accessFlags) {
		return this == other || other.IsSubClassOf(this)
	}
	if !IsPrivate(c.accessFlags) {
		return this.getPackageName() == other.getPackageName()
	}
	return this == other
}
