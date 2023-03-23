package intUtil

import "strconv"

// IntToString 将int类型转换为string类型
func IntToString(n int) string {
	return strconv.Itoa(n)
}

// Int8ToString 将int8类型转换为string类型
func Int8ToString(n int8) string {
	return strconv.Itoa(int(n))
}

// Int16ToString 将int16类型转换为string类型
func Int16ToString(n int16) string {
	return strconv.Itoa(int(n))
}

// Int32ToString 将int32类型转换为string类型
func Int32ToString(n int32) string {
	return strconv.Itoa(int(n))
}

// Int64ToString 将int64类型转换为string类型
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}
