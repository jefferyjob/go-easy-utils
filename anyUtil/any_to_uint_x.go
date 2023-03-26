package anyUtil

import (
	"github.com/jefferyjob/go-easy-utils"
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
		return 0, go_easy_utils.ErrValOut
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
		return 0, go_easy_utils.ErrValOut
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
		return 0, go_easy_utils.ErrValOut
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
		return 0, go_easy_utils.ErrValOut
	}
	return uint32(value), nil
}

// AnyToUint64 将给定的值转换为 uint64
func AnyToUint64(i interface{}) (uint64, error) {
	if i == nil {
		return 0, nil
	}

	// 检查解引用后的值是否为 nil
	if reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil() {
		return 0, nil
	}

	v := reflect.ValueOf(i)
	// 处理指针类型
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.ValueOf(i).IsNil() {
			return 0, nil
		}
		v = reflect.ValueOf(i).Elem()
	}

	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue := v.Int()
		if intValue < 0 {
			return 0, nil
		}
		return uint64(intValue), nil
	case reflect.Float32, reflect.Float64:
		floatValue := v.Float()
		if floatValue < 0 {
			return 0, nil
		}
		return uint64(floatValue), nil
	case reflect.Complex64, reflect.Complex128:
		realValue := real(v.Complex())
		if realValue < 0 {
			return 0, nil
		}
		return uint64(realValue), nil
	case reflect.String:
		strValue := v.String()
		uintValue, err := strconv.ParseUint(strValue, 10, 64)
		if err != nil {
			return 0, go_easy_utils.ErrSyntax
		}
		return uintValue, nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	//case reflect.Ptr, reflect.Interface:
	//	if v.IsNil() {
	//		return 0, nil
	//	}
	//	return toUint64(v.Elem().Interface())
	default:
		return 0, go_easy_utils.ErrType
	}
}
