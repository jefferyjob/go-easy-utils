package strx

import "unsafe"

// ToBytes 字符串转字节数组
func ToBytes(v string) []byte {
	return *(*[]byte)(unsafe.Pointer(&v))
}
