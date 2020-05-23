package utils

import (
	"bytes"
	"encoding/binary"
	"unsafe"
	"math"
)

var Endian binary.ByteOrder

// CheckEndian checks the byte order and determines the endiannes.
func CheckEndian() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
		case [2]byte{0xCD, 0xAB}:
			Endian = binary.LittleEndian
		case [2]byte{0xAB, 0xCD}:
			Endian = binary.BigEndian
		default:
			panic("Unknown native endian.")
	}
}

// Float64ToByte converts a float64 type number to a byte array.
func Float64ToByte(f float64) ([]byte, error) {
	if Endian == nil {
		CheckEndian()
	}

	var buf bytes.Buffer
	err := binary.Write(&buf, Endian, f)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// BytesToFloat64 converts bytes to a float64.
func BytesToFloat64(bytes []byte) float64 {
    bits := Endian.Uint64(bytes)
    float := math.Float64frombits(bits)
    return float
}
