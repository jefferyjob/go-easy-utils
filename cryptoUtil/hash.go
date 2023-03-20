// Package cryptoUtil 加密与解密工具
package cryptoUtil

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashSHA256 hash加密
func HashSHA256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}
