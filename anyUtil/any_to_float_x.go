package anyUtil

import (
	"fmt"
	"math"
	"strconv"
)

// AnyToFloat32 将给定的值转换为float32
func AnyToFloat32(value interface{}) (float32, error) {
	f64, err := AnyToFloat64(value)
	if err != nil {
		return 0, err
	}
	fmt.Println(-math.MaxFloat32, f64, math.MaxFloat32)
	if f64 < -math.MaxFloat32 || f64 > math.MaxFloat32 {
		return 0, ErrValOut
	}
	return float32(f64), nil
}

// AnyToFloat64 将给定的值转换为float64
func AnyToFloat64(v interface{}) (float64, error) {
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
			return 0, ErrSyntax
		}
		return v, nil
	}
	return 0, ErrType

	//switch reflect.TypeOf(v).Kind() {
	//case reflect.Float64:
	//	return v.(float64), nil
	//case reflect.Float32:
	//	return float64(v.(float32)), nil
	//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//	intVal := reflect.ValueOf(v).Int()
	//	if float64(intVal) > math.MaxFloat64 || float64(intVal) < -math.MaxFloat64 {
	//		return 0, ErrValOut
	//	}
	//	return float64(intVal), nil
	//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	//	uintVal := reflect.ValueOf(v).Uint()
	//	if uintVal > uint64(math.MaxUint) {
	//		return 0, ErrValOut
	//	}
	//	return float64(uintVal), nil
	//case reflect.String:
	//	return strconv.ParseFloat(v.(string), 64)
	//default:
	//	return 0, ErrType
	//}
}
