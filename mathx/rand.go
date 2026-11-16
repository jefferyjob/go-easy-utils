package mathx

import (
	"math/rand"
)

// Rand 生成随机整数
func Rand[T Numeric](min, max T) T {
	return T(rand.Intn(int(max)-int(min)+1) + int(min))
}
