package strUtil

import "strconv"

// StrToUint string转uint
func StrToUint(str string) uint {
	i, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0
	}
	return uint(i)
}

// StrToUint8 string转uint8
func StrToUint8(str string) uint8 {
	i, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return 0
	}
	return uint8(i)
}

// StrToUint16 string转uint16
func StrToUint16(str string) uint16 {
	i, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0
	}
	return uint16(i)
}

// StrToUint32 string转uint32
func StrToUint32(str string) uint32 {
	i, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(i)
}

// StrToUint64 string转uint64
func StrToUint64(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
