package jsonUtil

import (
	"reflect"
	"strconv"
)

func toFloat64(i any) (float64, error) {
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
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		floatValue, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, ErrSyntax
		}
		return floatValue, nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case complex64:
		return float64(real(v)), nil
	case complex128:
		return real(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	//case *float32:
	//	return float64(*v), nil
	//case *float64:
	//	return *v, nil
	//case *string:
	//	floatValue, err := strconv.ParseFloat(*v, 64)
	//	if err != nil {
	//		return 0, ErrSyntax
	//	}
	//	return floatValue, nil
	//case *uint:
	//	return float64(*v), nil
	//case *uint8:
	//	return float64(*v), nil
	//case *uint16:
	//	return float64(*v), nil
	//case *uint32:
	//	return float64(*v), nil
	//case *uint64:
	//	return float64(*v), nil
	//case *int:
	//	return float64(*v), nil
	//case *int8:
	//	return float64(*v), nil
	//case *int16:
	//	return float64(*v), nil
	//case *int32:
	//	return float64(*v), nil
	//case *int64:
	//	return float64(*v), nil
	//case *complex64:
	//	return float64(real(*v)), nil
	//case *complex128:
	//	return real(*v), nil
	//case *bool:
	//	if v != nil && *v {
	//		return 1, nil
	//	} else {
	//		return 0, nil
	//	}
	default:
		return 0, ErrType
	}
}

func toFloat64Reflect(i any) (float64, error) {
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
		return v.Float(), nil
	case reflect.String:
		if v.String() == "" {
			return 0, nil
		}
		floatValue, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return 0, ErrSyntax
		}
		return floatValue, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), nil
	case reflect.Complex64, reflect.Complex128:
		return real(v.Complex()), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	//case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice:
	//	return 0, nil
	default:
		return 0, ErrType
	}
}
