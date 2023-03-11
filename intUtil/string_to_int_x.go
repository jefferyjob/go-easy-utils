package intUtil

import "strconv"

// StrToInt 将string类型转换为int类型
func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// StrToInt8 将string类型转换为int8类型
func StrToInt8(s string) (int8, error) {
	n, err := strconv.Atoi(s)
	return int8(n), err
}

// StrToInt16 将string类型转换为int16类型
func StrToInt16(s string) (int16, error) {
	n, err := strconv.Atoi(s)
	return int16(n), err
}

// StrToInt32 将string类型转换为int32类型
func StrToInt32(s string) (int32, error) {
	n, err := strconv.Atoi(s)
	return int32(n), err
}

// StrToInt64 将string类型转换为int64类型
func StrToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
