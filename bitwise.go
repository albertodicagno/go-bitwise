package bitwise

import "encoding/binary"

// BytesToUint16 converts a big endian array of bytes to an array of unit16s
func BytesToUint16(bytes []byte) []uint16 {
values := make([]uint16, len(bytes)/2)

for i := range values {
values[i] = binary.BigEndian.Uint16(bytes[i*2 : (i+1)*2])
}
return values
}

func BytesToInt32(bytes []byte) []int32 {
values := make([]int32, len(bytes)/2)

for i := range values {
values[i] = int32(binary.BigEndian.Uint16(bytes[i*2 : (i+1)*2]))
}
return values
}

func Uint16ToBytes(values []uint16) []byte {
bytes := make([]byte, len(values)*2)

for i, value := range values {
binary.BigEndian.PutUint16(bytes[i*2:(i+1)*2], value)
}
return bytes
}

func BitAtPosition(value uint8, pos uint) uint8 {
return (value >> pos) & 0x01
}