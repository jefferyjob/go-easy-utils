package floatUtil

import "math"

// Float64Round 四舍五入
func Float64Round(input float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	rounded := math.Round(input * shift)
	return rounded / shift
}

// Float64Ceil 向上取整
func Float64Ceil(input float64) float64 {
	return math.Ceil(input*100) / 100
}
