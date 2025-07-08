package mathx

import (
	"math"
)

// Floor 对float数据向下取整
func Floor[T float32 | float64](num T) int {
	return int(math.Floor(float64(num)))
}
