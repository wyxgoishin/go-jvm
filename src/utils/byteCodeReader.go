package utils

import (
	"encoding/binary"
)

type ByteCodeReader struct {
	data []byte
	pc   int
}

func NewByteCodeReader(byteData []byte) *ByteCodeReader {
	return &ByteCodeReader{
		data: byteData,
		pc:   0,
	}
}

func (reader *ByteCodeReader) ReadUint8() uint8 {
	val := reader.data[reader.pc]
	reader.pc++
	return val
}

func (reader *ByteCodeReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}

func (reader *ByteCodeReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(reader.data[reader.pc:])
	reader.pc += 2
	return val
}

func (reader *ByteCodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}

func (reader *ByteCodeReader) ReadUint16s() []uint16 {
	count := reader.ReadUint16()
	ret := make([]uint16, count)
	for i := range ret {
		ret[i] = reader.ReadUint16()
	}
	return ret
}

func (reader *ByteCodeReader) ReadInt16s() []int16 {
	count := reader.ReadUint16()
	ret := make([]int16, count)
	for i := range ret {
		ret[i] = reader.ReadInt16()
	}
	return ret
}

func (reader *ByteCodeReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(reader.data[reader.pc:])
	reader.pc += 4
	return val
}

func (reader *ByteCodeReader) ReadInt32() int32 {
	return int32(reader.ReadUint32())
}

func (reader *ByteCodeReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(reader.data[reader.pc:])
	reader.pc += 8
	return val
}

func (reader *ByteCodeReader) ReadInt64() int64 {
	return int64(reader.ReadUint64())
}

func (reader *ByteCodeReader) ReadBytes(n uint32) []byte {
	val := reader.data[reader.pc : reader.pc+int(n)]
	reader.pc += int(n)
	return val
}

func (reader *ByteCodeReader) SkipPadding() {
	for reader.pc%4 != 0 {
		reader.ReadUint8()
	}
}

func (reader *ByteCodeReader) Reset(data []byte, pc int) {
	reader.data = data
	reader.pc = pc
}

func (reader *ByteCodeReader) PC() int {
	return reader.pc
}

func (reader *ByteCodeReader) SetPC(pc int) {
	reader.pc = pc
}
