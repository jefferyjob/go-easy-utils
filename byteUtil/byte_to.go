package byteUtil

import "unsafe"

// BytesToStr 字节数组转字符串
func BytesToStr(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
