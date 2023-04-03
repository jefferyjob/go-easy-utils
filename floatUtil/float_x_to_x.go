package floatUtil

import (
	"fmt"
	"strconv"
)

// Float32ToStr float32转字符串
func Float32ToStr(f float32) string {
	return fmt.Sprintf("%v", f)
}

// Float64ToStr float64转字符串
func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Float32ToFloat64 float32转float64
func Float32ToFloat64(f float32) float64 {
	return float64(f)
}

// Float64ToFloat32 float64转float32
func Float64ToFloat32(f float64) float32 {
	return float32(f)
}
