package heap

import (
	"go-jvm/src/classfile"
)

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPc   int // pc of first setence of try{}
	endPc     int // pc of first setence after try{}
	handlerPc int
	catchType *ClassRef
}

func newExceptionTable(entries []*classfile.ExceptionTableEntry, rtcp *RuntimeConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for idx, entry := range entries {
		handler := &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(entry.CatchType(), rtcp), // 0 means catch all Exception
		}
		table[idx] = handler
	}
	return table
}

func getCatchType(index uint16, rtcp *RuntimeConstantPool) *ClassRef {
	if index == 0 {
		return nil
	}
	return rtcp.GetConstant(uint(index)).(*ClassRef)
}

// ToDo: multi-catch, nested try-catch, finally, https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-3.html#jvms-3.12
func (e ExceptionTable) findExceptionHandler(class *Class, pc int) *ExceptionHandler {
	for _, handler := range e {
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler // catch all
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == class || catchClass.IsSuperClassOf(class) {
				return handler
			}
		}
	}
	return nil
}
