package utils

import "math"

const (
	FLOAT32_INF     = 0x7F800000
	FLOAT32_NEG_INF = 0xFF800000
	FLOAT32_NAN     = 0x7F800001
	MAX_CHAR        = 0xFFFF
	MIN_CHAR        = 0x0000
)

func IsNaN(val float32) bool {
	bits := math.Float32bits(val)
	return (bits<<9 != 0) && (bits<<1>>24 == 255)
}

func NaN() float32 {
	return math.Float32frombits(FLOAT32_NAN)
}

func IsInf(val float32, sign int) bool {
	bits := math.Float32bits(val)
	firstBit := bits >> 31
	return (sign == 0 || sign*int(firstBit) > 0) && bits<<9 == 0 && bits<<1>>24 == 255
}
