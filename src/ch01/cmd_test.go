package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func Test_parseCmd(t *testing.T) {
	Convey("HelpFlagMatchTest", t, func() {
		Convey("Matches_help", func() {
			os.Args = append(os.Args, "-help")
			So(parseCmd().helpFlag, ShouldBeTrue)
		})
		//Convey("Matches_h", func() {
		//	os.Args = append(os.Args, "-h")
		//	So(parseCmd().helpFlag, ShouldBeTrue)
		//})
	})
}
