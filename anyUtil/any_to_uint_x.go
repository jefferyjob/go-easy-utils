package anyUtil

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// AnyToUint 将给定的值转换为 uint
func AnyToUint(input interface{}) (uint, error) {
	v, err := AnyToUint64(input)
	if err != nil {
		return 0, err
	}

	// uint 兼容32位和64位系统
	if uint64(uint(v)) != v {
		return 0, errors.New("value out of range")
	}

	return uint(v), nil
}

// AnyToUint8 将给定的值转换为 uint8
func AnyToUint8(input interface{}) (uint8, error) {
	value, err := AnyToUint64(input)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint8 {
		return 0, fmt.Errorf("%d out of uint8 range", value)
	}
	return uint8(value), nil
}

// AnyToUint16 将给定的值转换为 uint16
func AnyToUint16(input interface{}) (uint16, error) {
	value, err := AnyToUint64(input)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint16 {
		return 0, fmt.Errorf("%d out of uint16 range", value)
	}
	return uint16(value), nil
}

// AnyToUint32 将给定的值转换为 uint32
func AnyToUint32(input interface{}) (uint32, error) {
	value, err := AnyToUint64(input)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint32 {
		return 0, fmt.Errorf("%d out of uint32 range", value)
	}
	return uint32(value), nil
}

// AnyToUint64 将给定的值转换为 uint64
func AnyToUint64(value interface{}) (uint64, error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(value).Uint(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := reflect.ValueOf(value).Int()
		if v < 0 {
			return 0, fmt.Errorf("value %v is negative and cannot be converted to unsigned integer", v)
		}
		return uint64(v), nil
	case reflect.Float32, reflect.Float64:
		v := reflect.ValueOf(value).Float()
		if v < 0 {
			return 0, fmt.Errorf("value %f is too large or too small to convert to uint", v)
		}
		return uint64(v), nil
	case reflect.String:
		val, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return val, nil
	}
	return 0, fmt.Errorf("unsupported type %T to uint", value)
}
