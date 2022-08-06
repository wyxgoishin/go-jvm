package reserved

import (
	"fmt"
	"go-jvm/src/ch11/instruction/base"
	"go-jvm/src/ch11/native"
	_ "go-jvm/src/ch11/native/java/lang"
	_ "go-jvm/src/ch11/native/sun/misc"
	"go-jvm/src/ch11/rtda"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (inst *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		panic(fmt.Sprintf("java.lang.UnsatisfiedLinkError: %v.%v%v", className, methodName, methodDescriptor))
	}
	nativeMethod(frame)
}
