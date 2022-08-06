package io

import (
	"go-jvm/src/ch11/native"
	"go-jvm/src/ch11/rtda"
)

func init() {
	native.Register("java/io/FileInputStream", "initIDs", "()V", initIDs)
}

// private static native void initIDs();
// https://github.com/openjdk/blob/jdk8-b120/jdk/src/share/native/java/io/FileInputStream.c
func initIDs(frame *rtda.Frame) {

}
