package io

import (
	"go-jvm/src/native"
	"go-jvm/src/rtda"
)

func init() {
	native.Register("java/io/FileDescriptor", "initIDs", "()V", fdInitIDs)
}

// private static native void initIDs();
// ToDo: implement me
func fdInitIDs(frame *rtda.Frame) {

}
