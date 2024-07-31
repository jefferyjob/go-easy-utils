package jsonUtil

import (
	"reflect"
	"strconv"
)

func toUint64Reflect(i any) (uint64, error) {
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
		if v.String() == "" {
			return 0, nil
		}
		uintValue, err := strconv.ParseUint(v.String(), 10, 64)
		if err != nil {
			return 0, ErrSyntax
		}
		return uintValue, nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, ErrType
	}
}
