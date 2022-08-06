package io

import (
	"go-jvm/src/ch11/native"
	"go-jvm/src/ch11/rtda"
	"os"
	"unsafe"
)

const fos = "java/io/FileOutputStream"

func init() {
	native.Register(fos, "writeBytes", "([BIIZ)V", writeBytes)
}

// private native void writeBytes(byte b[], int off, int len, boolean append) throws IOException;
// ([BIIZ)V
func writeBytes(frame *rtda.Frame) {
	vars := frame.LocalVars
	//this := vars.GetRef(0)
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	byteLen := vars.GetInt(3)
	//append := vars.GetBoolean(4)

	jvBytes := b.Data().([]int8)
	goBytes := castInt8sToUint8s(jvBytes)
	goBytes = goBytes[off : off+byteLen]
	os.Stdout.Write(goBytes)
}

func castInt8sToUint8s(jBytes []int8) (goBytes []byte) {
	ptr := unsafe.Pointer(&jBytes)
	goBytes = *((*[]byte)(ptr))
	return
}
