package anyUtil

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

// AnyToInt 将给定的值转换为 int
func AnyToInt(input interface{}) (int, error) {
	v, err := AnyToInt64(input)
	if err != nil {
		return 0, err
	}
	if v < int64(^int(0)+1) || v > int64(^int(0)) {
		return 0, errors.New("value out of range")
	}
	return int(v), nil
}

// AnyToInt8 将给定的值转换为 int8
func AnyToInt8(input interface{}) (int8, error) {
	value, err := AnyToInt64(input)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt8 || value > math.MaxInt8 {
		return 0, fmt.Errorf("%d out of int8 range", value)
	}
	return int8(value), nil
}

// AnyToInt16 将给定的值转换为 int16
func AnyToInt16(input interface{}) (int16, error) {
	value, err := AnyToInt64(input)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt16 || value > math.MaxInt16 {
		return 0, fmt.Errorf("%d out of int16 range", value)
	}
	return int16(value), nil
}

// AnyToInt32 将给定的值转换为 int32
func AnyToInt32(input interface{}) (int32, error) {
	value, err := AnyToInt64(input)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt32 || value > math.MaxInt32 {
		return 0, fmt.Errorf("%d out of int32 range", value)
	}
	return int32(value), nil
}

// AnyToInt64 将给定的值转换为 int64
func AnyToInt64(value interface{}) (int64, error) {
	//switch input.(type) {
	//case int:
	//	return int64(input.(int)), nil
	//case int8:
	//	return int64(input.(int8)), nil
	//case int16:
	//	return int64(input.(int16)), nil
	//case int32:
	//	return int64(input.(int32)), nil
	//case int64:
	//	return input.(int64), nil
	//case uint:
	//	return int64(input.(uint)), nil
	//case uint8:
	//	return int64(input.(uint8)), nil
	//case uint16:
	//	return int64(input.(uint16)), nil
	//case uint32:
	//	return int64(input.(uint32)), nil
	//case uint64:
	//	return int64(input.(uint64)), nil
	//case string:
	//	val, err := strconv.ParseInt(input.(string), 10, 64)
	//	if err != nil {
	//		return 0, err
	//	}
	//	return val, nil
	//case float32:
	//	return int64(input.(float32)), nil
	//case float64:
	//	return int64(input.(float64)), nil
	//default:
	//	return 0, fmt.Errorf("unsupported type %T", input)
	//}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(value).Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := reflect.ValueOf(value).Uint()
		if v > uint64(^int64(0)) {
			return 0, errors.New("value out of range")
		}
		return int64(v), nil
	case reflect.Float32, reflect.Float64:
		v := reflect.ValueOf(value).Float()
		if v < float64(^int64(0))+1 || v > float64(^int64(0)) {
			return 0, errors.New("value out of range")
		}
		return int64(v), nil
	}
	return 0, errors.New("unsupported type")
}
