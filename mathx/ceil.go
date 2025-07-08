package mathx

import (
	"math"
)

// Ceil 对float数据向上取整
func Ceil[T float32 | float64](num T) int {
	return int(math.Ceil(float64(num)))
}
