package jsonUtil

import (
	"github.com/jefferyjob/go-easy-utils/v2"
	"reflect"
	"strconv"
)

func toUint64(i any) (uint64, error) {
	if i == nil {
		return 0, nil
	}

	// 处理指针类型
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.ValueOf(i).IsNil() {
			return 0, nil
		}
		i = reflect.ValueOf(i).Elem().Interface()
	}

	switch v := i.(type) {
	case float32:
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case float64:
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case complex64:
		if real(v) < 0 || imag(v) < 0 {
			return 0, nil
		}
		return uint64(real(v)), nil
	case complex128:
		if real(v) < 0 || imag(v) < 0 {
			return 0, nil
		}
		return uint64(real(v)), nil
	case string:
		intValue, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, go_easy_utils.ErrSyntax
		}
		return intValue, nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case int:
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	//case *float32:
	//	if *v < 0 {
	//		return 0, nil
	//	}
	//	return uint64(*v), nil
	//case *float64:
	//	if *v < 0 {
	//		return 0, nil
	//	}
	//	return uint64(*v), nil
	//case *complex64:
	//	if real(*v) < 0 || imag(*v) < 0 {
	//		return 0, nil
	//	}
	//	return uint64(real(*v)), nil
	//case *complex128:
	//	if real(*v) < 0 || imag(*v) < 0 {
	//		return 0, nil
	//	}
	//	return uint64(real(*v)), nil
	//case *string:
	//	intValue, err := strconv.ParseUint(*v, 10, 64)
	//	if err != nil {
	//		return 0, go_easy_utils.ErrSyntax
	//	}
	//	return intValue, nil
	//case *bool:
	//	if *v {
	//		return 1, nil
	//	} else {
	//		return 0, nil
	//	}
	//case *uint:
	//	return uint64(*v), nil
	//case *uint8:
	//	return uint64(*v), nil
	//case *uint16:
	//	return uint64(*v), nil
	//case *uint32:
	//	return uint64(*v), nil
	//case *uint64:
	//	return *v, nil
	//case *int:
	//	if *v < 0 {
	//		return 0, nil
	//	}
	//	return uint64(*v), nil
	//case *int8:
	//	if *v < 0 {
	//		return 0, nil
	//	}
	//	return uint64(*v), nil
	//case *int16:
	//	if *v < 0 {
	//		return 0, nil
	//	}
	//	return uint64(*v), nil
	//case *int32:
	//	if *v < 0 {
	//		return 0, nil
	//	}
	//	return uint64(*v), nil
	//case *int64:
	//	if *v < 0 {
	//		return 0, nil
	//	}
	//	return uint64(*v), nil
	default:
		return 0, go_easy_utils.ErrType
	}
}

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
			return 0, go_easy_utils.ErrSyntax
		}
		return uintValue, nil
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
