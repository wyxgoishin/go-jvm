package lang

import (
	"go-jvm/src/native"
	"go-jvm/src/rtda"
	"go-jvm/src/rtda/heap"
)

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy)
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)
	stackTraceElements := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stackTraceElements)
}

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func createStackTraceElements(obj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	skip := distanceToObject(obj.Class()) + 2 // 2 -> fillInStackTrace(int) 和 fillInStackTrace()方法的栈帧
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JvName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}

// 计算需要跳过的异常类的初始化栈帧数量
func distanceToObject(class *heap.Class) int {
	dist := -1
	for ; class != nil; class = class.SuperClass() {
		dist++
	}
	return dist
}
