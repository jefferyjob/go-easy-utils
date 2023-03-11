package randUtil

import (
	"math/rand"
	"time"
)

// RandomInt 生成指定范围内的随机整数
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// RandomFloat 生成指定范围内的随机浮点数
func RandomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}
