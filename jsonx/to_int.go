package jsonx

import (
	"reflect"
	"strconv"
)

func toInt64Reflect(i any) (int64, error) {
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
		if v.String() == "" {
			return 0, nil
		}
		intValue, err := strconv.ParseInt(v.String(), 10, 64)
		if err != nil {
			return 0, ErrSyntax
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
		return 0, ErrType
	}
}
