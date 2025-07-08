package strx

import (
	"strconv"
)

// ToInt string转int
func ToInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return i
}

// ToInt8 string转int8
func ToInt8(v string) int8 {
	i, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return 0
	}
	return int8(i)
}

// ToInt16 string转int16
func ToInt16(v string) int16 {
	i, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		return 0
	}
	return int16(i)
}

// ToInt32 string转int32
func ToInt32(v string) int32 {
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

// ToInt64 string转int64
func ToInt64(v string) int64 {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
