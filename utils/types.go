package utils

import (
	"bytes"
	"encoding/binary"
	"math"
	"unsafe"
)

var endian binary.ByteOrder

// CheckEndian checks the byte order and determines the endiannes.
func CheckEndian() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		endian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		endian = binary.BigEndian
	default:
		panic("Unknown native endian.")
	}
}

// Abs returns the absolute value of x.
func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}

	return x
}

// Float64ToByte converts a float64 type number to a byte array.
func Float64ToByte(f float64) ([]byte, error) {
	if endian == nil {
		CheckEndian()
	}

	var buf bytes.Buffer
	err := binary.Write(&buf, endian, f)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// BytesToFloat64 converts bytes to a float64.
func BytesToFloat64(bytes []byte) float64 {
	if endian == nil {
		CheckEndian()
	}

	bits := endian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

// UInt64ToBytes converts an uint64 to []bytes.
func UInt64ToBytes(num uint64) []byte {
	if endian == nil {
		CheckEndian()
	}

	bytes := make([]byte, 8)
	endian.PutUint64(bytes, num)

	return bytes
}

// BytesToUInt64 converts a byte array to an int64.
func BytesToUInt64(bytes []byte) uint64 {
	if endian == nil {
		CheckEndian()
	}

	return endian.Uint64(bytes)
}
