package format

import "fmt"

var ErrorInvalidLength = fmt.Errorf("invalid data length")

func Uint32AtOffset(data []byte, offset int) uint32 {
	return uint32(data[offset]) | uint32(data[offset+1])<<8 | uint32(data[offset+2])<<16 | uint32(data[offset+3])<<24
}
