package bytex

import "unsafe"

// ToStr 字节数组转字符串
func ToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
