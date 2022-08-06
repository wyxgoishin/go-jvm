package rtda

import (
	. "github.com/smartystreets/goconvey/convey"
	"go-jvm/src/ch06/rtda/heap"
	"testing"
)

func TestLocalVars(t *testing.T) {
	Convey("TestLocalVars", t, func() {
		localVars := newLocalVars(10)

		localVars.SetInt(0, 10)
		localVars.SetLong(1, 3)
		localVars.SetFloat(3, 2.4e-2)
		localVars.SetDouble(4, 9e2)
		localVars.SetRef(6, &heap.Object{})

		So(localVars.GetInt(0), ShouldEqual, 10)
		So(localVars.GetLong(1), ShouldEqual, 3)
		So(localVars.GetFloat(3), ShouldEqual, 2.4e-2)
		So(localVars.GetDouble(4), ShouldEqual, 9e2)
		So(localVars.GetRef(6), ShouldNotBeNil)
	})
}
