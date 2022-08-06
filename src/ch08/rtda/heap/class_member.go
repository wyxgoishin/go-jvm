package heap

import "go-jvm/src/ch08/classfile"

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
	if c.IsPublic() {
		return true
	}
	this := c.class
	if this.IsProtected() {
		return this == other || other.IsSubClassOf(this)
	}
	if c.IsPrivate() {
		return this.GetPackageName() == other.GetPackageName()
	}
	return this == other
}

func (c *ClassMember) AccessFlags() uint16 {
	return c.accessFlags
}

func (c *ClassMember) Name() string {
	return c.name
}

func (c *ClassMember) Descriptor() string {
	return c.descriptor
}

func (c *ClassMember) IsPublic() bool {
	return c.accessFlags&ACC_PUBLIC > 0
}

func (c *ClassMember) IsPrivate() bool {
	return c.accessFlags&ACC_PRIVATE > 0
}

func (c *ClassMember) IsProtected() bool {
	return c.accessFlags&ACC_PROTECTED > 0
}

func (c *ClassMember) IsStatic() bool {
	return c.accessFlags&ACC_STATIC > 0
}

func (c *ClassMember) IsFinal() bool {
	return c.accessFlags&ACC_FINAL > 0
}

func (c *ClassMember) IsSynthetic() bool {
	return c.accessFlags&ACC_SYNTHETIC > 0
}
