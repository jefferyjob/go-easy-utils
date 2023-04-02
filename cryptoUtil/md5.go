package cryptoUtil

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 MD5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}
