package randUtil

import (
	"math/rand"
	"time"
)

// RandomString 生成指定长度的随机字符串
func RandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomBytes 生成指定长度的随机字节数组
func RandomBytes(length int) []byte {
	b := make([]byte, length)
	rand.Read(b)
	return b
}
