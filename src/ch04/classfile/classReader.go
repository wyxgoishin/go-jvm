package classfile

import (
	"encoding/binary"
	"fmt"
)

type ClassReader struct {
	data []byte
}

func newClassReader(classData []byte) *ClassReader {
	return &ClassReader{
		data: classData,
	}
}

func (r *ClassReader) readUint8() uint8 {
	val := r.data[0]
	r.data = r.data[1:]
	return val
}

func (r *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(r.data)
	r.data = r.data[2:]
	return val
}

func (r *ClassReader) readUint16s() []uint16 {
	count := r.readUint16()
	ret := make([]uint16, count)
	for i := range ret {
		ret[i] = r.readUint16()
	}
	return ret
}

func (r *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(r.data)
	r.data = r.data[4:]
	return val
}

func (r *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(r.data)
	r.data = r.data[8:]
	return val
}

func (r *ClassReader) readBytes(n uint32) []byte {
	if int(n) > len(r.data) {
		panic(fmt.Sprintf("read num %q exceeds remained num of %q", n, len(r.data)))
	}
	val := r.data[:n]
	r.data = r.data[n:]
	return val
}
