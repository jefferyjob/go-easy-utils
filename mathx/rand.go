package mathx

import (
	"math/rand"
	"time"
)

// Rand 生成随机整数
func Rand[T Numeric](min, max T) T {
	rand.Seed(time.Now().UnixNano())
	return T(rand.Intn(int(max)-int(min)+1) + int(min))
}
