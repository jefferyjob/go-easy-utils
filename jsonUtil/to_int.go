package jsonUtil

import (
	"reflect"
	"strconv"
)

func toInt64(i any) (int64, error) {
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
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case complex64:
		return int64(real(v)), nil
	case complex128:
		return int64(real(v)), nil
	case string:
		intValue, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, ErrSyntax
		}
		return intValue, nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	//case *bool:
	//	if v == nil {
	//		return 0, nil
	//	}
	//	if *v {
	//		return 1, nil
	//	}
	//	return 0, nil
	//case *uint:
	//	return int64(*v), nil
	//case *uint8:
	//	return int64(*v), nil
	//case *uint16:
	//	return int64(*v), nil
	//case *uint32:
	//	return int64(*v), nil
	//case *uint64:
	//	return int64(*v), nil
	//case *float32:
	//	return int64(*v), nil
	//case *float64:
	//	return int64(*v), nil
	//case *complex64:
	//	if v == nil {
	//		return 0, nil
	//	}
	//	return int64(real(*v)), nil
	//case *complex128:
	//	if v == nil {
	//		return 0, nil
	//	}
	//	return int64(real(*v)), nil
	//case *string:
	//	if v == nil {
	//		return 0, nil
	//	}
	//	intValue, err := strconv.ParseInt(*v, 10, 64)
	//	if err != nil {
	//		return 0, ErrSyntax
	//	}
	//	return intValue, nil
	//case *int:
	//	return int64(*v), nil
	//case *int8:
	//	return int64(*v), nil
	//case *int16:
	//	return int64(*v), nil
	//case *int32:
	//	return int64(*v), nil
	//case *int64:
	//	return *v, nil
	default:
		return 0, ErrType
	}
}

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
