package reference

import (
	"fmt"
	"go-jvm/src/ch10/instruction/base"
	"go-jvm/src/ch10/rtda"
	"go-jvm/src/ch10/rtda/heap"
	"reflect"
)

type ATHROW struct {
	base.NoOperandsInstruction
}

func (inst *ATHROW) Execute(frame *rtda.Frame) {
	except := frame.OperandStack.PopRef()
	if except == nil {
		panic("java.lang.NullPointerException")
	}
	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, except) {
		handleUncaughtException(thread, except)
	}
}

func findAndGotoExceptionHandler(thread *rtda.Thread, except *heap.Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1
		handlerPC := frame.Method().FindExceptionHandler(except.Class(), pc)
		if handlerPC > 0 {
			frame.OperandStack.Clear()
			frame.OperandStack.PushRef(except)
			frame.SetNextPC(handlerPC)
			return true
		}
		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

func handleUncaughtException(thread *rtda.Thread, except *heap.Object) {
	thread.ClearStack()
	jvMsg := except.GetInstanceRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jvMsg)
	fmt.Printf("%v: %v", jvMsg.Class().JvName(), goMsg)
	stackInfos := reflect.ValueOf(except.Extra())
	for i := 0; i < stackInfos.Len(); i++ {
		info := stackInfos.Index(i).Interface().(interface{ String() string })
		fmt.Printf("\tat %v", info.String())
	}
}
