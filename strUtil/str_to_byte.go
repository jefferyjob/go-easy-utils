package strUtil

import "unsafe"

// StrToBytes 字符串转字节数组
func StrToBytes(v string) []byte {
	return *(*[]byte)(unsafe.Pointer(&v))
}
