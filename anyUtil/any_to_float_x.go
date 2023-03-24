package anyUtil

import (
	"github.com/jefferyjob/go-easy-utils"
	"math"
	"strconv"
)

// AnyToFloat32 将给定的值转换为float32
func AnyToFloat32(value interface{}) (float32, error) {
	f64, err := AnyToFloat64(value)
	if err != nil {
		return 0, err
	}
	if f64 < -math.MaxFloat32 || f64 > math.MaxFloat32 {
		return 0, go_easy_utils.ErrValOut
	}
	return float32(f64), nil
}

// AnyToFloat64 将给定的值转换为float64
func AnyToFloat64(v interface{}) (float64, error) {
	if v == nil {
		return 0, nil
	}

	switch val := v.(type) {
	case float32:
		return float64(val), nil
	case float64:
		return val, nil
	case int:
		return float64(val), nil
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case uint:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case string:
		v, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, go_easy_utils.ErrSyntax
		}
		return v, nil
	case interface{}:
		return 0, nil
	}
	return 0, go_easy_utils.ErrType
}
