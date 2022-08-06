package heap

import (
	"fmt"
	"go-jvm/src/ch10/classfile"
)

type Constant interface {
}

type RuntimeConstantPool struct {
	class  *Class
	consts []Constant
}

func newRuntimeConstantPool(class *Class, cfcp classfile.ConstantPool) *RuntimeConstantPool {
	cpCount := len(cfcp)
	consts := make([]Constant, cpCount)
	rtcp := &RuntimeConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfcp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantStringInfo:
			strInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = strInfo.String()
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtcp, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldRefInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(rtcp, fieldRefInfo)
		case *classfile.ConstantMethodRefInfo:
			methodRefInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rtcp, methodRefInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			interfaceMethodInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			consts[i] = newInterfaceMethodRef(rtcp, interfaceMethodInfo)
		}
	}
	return rtcp
}

func (rtcp *RuntimeConstantPool) GetConstant(index uint) Constant {
	if val := rtcp.consts[index]; val != nil {
		return val
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
