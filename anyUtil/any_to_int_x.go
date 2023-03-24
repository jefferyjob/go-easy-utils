package anyUtil

import (
	"fmt"
	"github.com/jefferyjob/go-easy-utils"
	"math"
	"reflect"
	"strconv"
)

// AnyToInt 将给定的值转换为 int
func AnyToInt(input interface{}) (int, error) {
	v, err := AnyToInt64(input)
	if err != nil {
		return 0, err
	}

	// int 兼容32位和64位系统
	if int64(int(v)) != v {
		fmt.Println(int64(int(v)), v)
		return 0, go_easy_utils.ErrValOut
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
		return 0, go_easy_utils.ErrValOut
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
		return 0, go_easy_utils.ErrValOut
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
		return 0, go_easy_utils.ErrValOut
	}
	return int32(value), nil
}

// AnyToInt64 将给定的值转换为 int64
func AnyToInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, nil
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(value).Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := reflect.ValueOf(value).Uint()
		if v > math.MaxUint64 {
			return 0, go_easy_utils.ErrValOut
		}
		return int64(v), nil
	case reflect.Float32, reflect.Float64:
		v := reflect.ValueOf(value).Float()
		if v < float64(math.MinInt64) || v > float64(math.MaxInt64) {
			return 0, go_easy_utils.ErrValOut
		}
		return int64(v), nil
	case reflect.String:
		val, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return 0, go_easy_utils.ErrSyntax
		}
		return val, nil
	case reflect.Interface:
		return AnyToInt64(reflect.ValueOf(value).Elem().Interface())
	}
	return 0, go_easy_utils.ErrType
}
