package mathx

import "math"

// Abs 返回一个数的绝对值
func Abs[T Numeric](num T) T {
	return T(math.Abs(float64(num)))
}
