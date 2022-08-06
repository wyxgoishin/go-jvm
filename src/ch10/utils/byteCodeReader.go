package utils

import (
	"encoding/binary"
)

type ByteCodeReader struct {
	data []byte
	pc   int
}

func NewByteCodeReader(classData []byte) *ByteCodeReader {
	return &ByteCodeReader{
		data: classData,
		pc:   0,
	}
}

func (r *ByteCodeReader) ReadUint8() uint8 {
	val := r.data[r.pc]
	r.pc++
	return val
}

func (r *ByteCodeReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(r.data[r.pc:])
	r.pc += 2
	return val
}

func (r *ByteCodeReader) ReadUint16s() []uint16 {
	count := r.ReadUint16()
	ret := make([]uint16, count)
	for i := range ret {
		ret[i] = r.ReadUint16()
	}
	return ret
}

func (r *ByteCodeReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(r.data[r.pc:])
	r.pc += 4
	return val
}

func (r *ByteCodeReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(r.data[r.pc:])
	r.pc += 8
	return val
}

func (r *ByteCodeReader) ReadBytes(n uint32) []byte {
	val := r.data[r.pc : r.pc+int(n)]
	r.pc += int(n)
	return val
}

func (r *ByteCodeReader) SkipPadding() {
	for r.pc%4 != 0 {
		r.ReadUint8()
	}
}

func (r *ByteCodeReader) Reset(data []byte, pc int) {
	r.data = data
	r.pc = pc
}

func (r *ByteCodeReader) PC() int {
	return r.pc
}

func (r *ByteCodeReader) SetPC(pc int) {
	r.pc = pc
}
