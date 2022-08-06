package rtda

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOperandStack(t *testing.T) {
	Convey("TestOperandStack", t, func() {
		operandStack := newOperandStack(10)

		var ival int32 = 10
		operandStack.PushInt(ival)
		So(operandStack.size, ShouldEqual, 1)
		So(operandStack.PopInt(), ShouldEqual, ival)

		var fval float32 = 10.0
		operandStack.PushFloat(fval)
		So(operandStack.size, ShouldEqual, 1)
		So(operandStack.PopFloat(), ShouldEqual, fval)

		var lval int64 = 100
		operandStack.PushLong(lval)
		So(operandStack.size, ShouldEqual, 2)
		So(operandStack.PopLong(), ShouldEqual, lval)

		var dval float64 = 3.4e9
		operandStack.PushDouble(dval)
		So(operandStack.size, ShouldEqual, 2)
		So(operandStack.PopDouble(), ShouldEqual, dval)

		ref := &Object{}
		operandStack.PushRef(ref)
		So(operandStack.size, ShouldEqual, 1)
		So(operandStack.PopRef(), ShouldEqual, ref)
	})
}
