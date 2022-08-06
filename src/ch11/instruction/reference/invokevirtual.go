package reference

import (
	"fmt"
	"go-jvm/src/ch11/instruction/base"
	"go-jvm/src/ch11/rtda"
	"go-jvm/src/ch11/rtda/heap"
	"strings"
)

type INVOKE_VIRTUAL struct {
	base.Index16Struction
}

//ToDo: currently only hack
func (i *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	curClass := frame.Method().Class()
	rtcp := curClass.RuntimeConstantPool()
	methodRef := rtcp.GetConstant(i.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack.GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		// hack
		if strings.HasPrefix(methodRef.Name(), "print") {
			_println(frame, methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerException")
	}
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(curClass) && resolvedMethod.Class().GetPackageName() != curClass.GetPackageName() && ref.Class() != curClass && !ref.Class().IsSubClassOf(curClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), resolvedMethod.Name(), resolvedMethod.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractmethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

func _println(frame *rtda.Frame, descriptor string) {
	stack := frame.OperandStack
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		jvStr := stack.PopRef()
		goStr := heap.GoString(jvStr)
		fmt.Printf("%v\n", goStr)
	case "(Ljava/lang/Object;)V":
		obj := stack.PopRef()
		slots := obj.Slots()
		for _, slot := range slots {
			slot.Print(true)
		}
		fmt.Println()
	default:
		fmt.Printf("println: %v\n", descriptor)
	}
	stack.PopRef()
}
