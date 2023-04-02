package strUtil

import (
	"strconv"
)

// StrToInt string转int
func StrToInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return i
}

// StrToInt8 string转int8
func StrToInt8(v string) int8 {
	i, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return 0
	}
	return int8(i)
}

// StrToInt16 string转int16
func StrToInt16(v string) int16 {
	i, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		return 0
	}
	return int16(i)
}

// StrToInt32 string转int32
func StrToInt32(v string) int32 {
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

// StrToInt64 string转int64
func StrToInt64(v string) int64 {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
