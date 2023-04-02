package byteUtil

import "unsafe"

// BytesToStr 字节数组转字符串
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
