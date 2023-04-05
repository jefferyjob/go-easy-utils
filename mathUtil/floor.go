package mathUtil

import (
	"math"
	"reflect"
)

// Floor 对float数据向下取整
func Floor[T float32 | float64](num T) int {
	switch reflect.ValueOf(num).Kind() {
	case reflect.Float32:
		return int(math.Floor(float64(num)))
	case reflect.Float64:
		return int(math.Floor(float64(num)))
	}
	return 0
}
