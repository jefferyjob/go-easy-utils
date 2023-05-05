package mathUtil

import (
	"math"
)

// Round 对float数据四舍五入
func Round[T float32 | float64](num T) int {
	return int(math.Round(float64(num)))
}
