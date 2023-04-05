package mathUtil

import (
	"math"
	"reflect"
)

// Round 对float数据四舍五入
func Round[T float32 | float64](num T) int {
	switch reflect.ValueOf(num).Kind() {
	case reflect.Float32:
		return int(math.Round(float64(num)))
	case reflect.Float64:
		return int(math.Round(float64(num)))
	}
	return 0
}
