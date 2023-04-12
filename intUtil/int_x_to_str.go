package intUtil

import "strconv"

// IntToStr 将int类型转换为string类型
func IntToStr(v int) string {
	return strconv.Itoa(v)
}

// Int8ToStr 将int8类型转换为string类型
func Int8ToStr(v int8) string {
	return strconv.Itoa(int(v))
}

// Int16ToString 将int16类型转换为string类型
func Int16ToStr(v int16) string {
	return strconv.Itoa(int(v))
}

// Int32ToStr 将int32类型转换为string类型
func Int32ToStr(v int32) string {
	return strconv.Itoa(int(v))
}

// Int64ToStr 将int64类型转换为string类型
func Int64ToStr(v int64) string {
	return strconv.FormatInt(v, 10)
}
