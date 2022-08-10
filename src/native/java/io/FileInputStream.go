package io

import (
	"go-jvm/src/native"
	"go-jvm/src/rtda"
)

func init() {
	native.Register("java/io/FileInputStream", "initIDs", "()V", fisInitIDs)
}

// private static native void initIDs();
// ToDo: implement me according to https://github.com/openjdk/jdk/blob/jdk8-b120/jdk/src/share/native/java/io/FileInputStream.c
func fisInitIDs(frame *rtda.Frame) {

}
