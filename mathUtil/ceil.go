package mathUtil

import (
	"math"
	"reflect"
)

// Ceil 对float数据向上取整
func Ceil[T float32 | float64](num T) int {
	switch reflect.ValueOf(num).Kind() {
	case reflect.Float32:
		return int(math.Ceil(float64(num)))
	case reflect.Float64:
		return int(math.Ceil(float64(num)))
	}
	return 0
}
