package reference

import (
	"go-jvm/src/ch07/instruction/base"
	"go-jvm/src/ch07/rtda"
	"go-jvm/src/ch07/rtda/heap"
)

type INVOKE_SPECIAL struct {
	base.Index16Struction
}

func (i *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	curClass := frame.Method().Class()
	rtcp := curClass.RuntimeConstantPool()
	methodRef := rtcp.GetConstant(i.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack.GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	if resolvedMethod.IsProtected() && curClass.IsSubClassOf(resolvedMethod.Class()) && resolvedMethod.Class().GetPackageName() != curClass.GetPackageName() && ref.Class() != curClass && !curClass.IsSubClassOf(ref.Class()) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := resolvedMethod
	if curClass.IsSuper() && resolvedClass.IsSuperClassOf(curClass) && resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(curClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
