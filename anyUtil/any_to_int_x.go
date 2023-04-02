package anyUtil

import (
	"github.com/jefferyjob/go-easy-utils"
	"math"
	"reflect"
	"strconv"
)

// AnyToInt 将给定的值转换为 int
func AnyToInt(i interface{}) (int, error) {
	v, err := AnyToInt64(i)
	if err != nil {
		return 0, err
	}

	// int 兼容32位和64位系统
	if int64(int(v)) != v {
		return 0, go_easy_utils.ErrValOut
	}

	return int(v), nil
}

// AnyToInt8 将给定的值转换为 int8
func AnyToInt8(i interface{}) (int8, error) {
	value, err := AnyToInt64(i)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt8 || value > math.MaxInt8 {
		return 0, go_easy_utils.ErrValOut
	}
	return int8(value), nil
}

// AnyToInt16 将给定的值转换为 int16
func AnyToInt16(i interface{}) (int16, error) {
	value, err := AnyToInt64(i)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt16 || value > math.MaxInt16 {
		return 0, go_easy_utils.ErrValOut
	}
	return int16(value), nil
}

// AnyToInt32 将给定的值转换为 int32
func AnyToInt32(i interface{}) (int32, error) {
	value, err := AnyToInt64(i)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt32 || value > math.MaxInt32 {
		return 0, go_easy_utils.ErrValOut
	}
	return int32(value), nil
}

// AnyToInt64 将给定的值转换为 int64
func AnyToInt64(i interface{}) (int64, error) {
	if i == nil {
		return 0, nil
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return 0, nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return int64(v.Float()), nil
	case reflect.String:
		intValue, err := strconv.ParseInt(v.String(), 10, 64)
		if err != nil {
			return 0, go_easy_utils.ErrSyntax
		}
		return intValue, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int64(v.Uint()), nil
	case reflect.Complex64:
		return int64(real(v.Complex())), nil
	case reflect.Complex128:
		return int64(real(v.Complex())), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, go_easy_utils.ErrType
	}
}
