package classfile

import (
	"fmt"
	"go-jvm/src/ch07/utils"
	"unicode/utf16"
)

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8Info struct {
	str string
}

func (c *ConstantUtf8Info) readInfo(cr *utils.ByteCodeReader) {
	length := uint32(cr.ReadUint16())
	bytes := cr.ReadBytes(length)
	c.str = decodeMUTF8(bytes)
}

func (c *ConstantUtf8Info) String() string {
	return c.str
}

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(byteArr []byte) string {
	utfLen := len(byteArr)
	charArr := make([]uint16, utfLen)

	var c, char2, char3 uint16
	count := 0
	charArrLen := 0

	for count < utfLen {
		c = uint16(byteArr[count])
		if c > 127 {
			break
		}
		count++
		charArr[charArrLen] = c
		charArrLen++
	}

	for count < utfLen {
		c = uint16(byteArr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			charArr[charArrLen] = c
			charArrLen++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(byteArr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			charArr[charArrLen] = c&0x1F<<6 | char2&0x3F
			charArrLen++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(byteArr[count-2])
			char3 = uint16(byteArr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			charArr[charArrLen] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			charArrLen++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utfLen
	charArr = charArr[0:charArrLen]
	runes := utf16.Decode(charArr)
	return string(runes)
}
