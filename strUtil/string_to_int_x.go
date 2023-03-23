package strUtil

import (
	"strconv"
	"unsafe"
)

// StrToInt string转int
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// StrToInt8 string转int8
func StrToInt8(str string) int8 {
	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0
	}
	return int8(i)
}

// StrToInt16 string转int16
func StrToInt16(str string) int16 {
	i, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0
	}
	return int16(i)
}

// StrToInt32 string转int32
func StrToInt32(str string) int32 {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

// StrToInt64 string转int64
func StrToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

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

// StrToBytes 字符串转字节数组
func StrToBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}
