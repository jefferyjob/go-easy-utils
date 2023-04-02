package intUtil

import "strconv"

// IntToString 将int类型转换为string类型
func IntToString(v int) string {
	return strconv.Itoa(v)
}

// Int8ToString 将int8类型转换为string类型
func Int8ToString(v int8) string {
	return strconv.Itoa(int(v))
}

// Int16ToString 将int16类型转换为string类型
func Int16ToString(v int16) string {
	return strconv.Itoa(int(v))
}

// Int32ToString 将int32类型转换为string类型
func Int32ToString(v int32) string {
	return strconv.Itoa(int(v))
}

// Int64ToString 将int64类型转换为string类型
func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}
