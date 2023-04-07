package jsonUtil

import (
	"fmt"
	"reflect"
	"strconv"
)

func toString(i any) string {
	if i == nil {
		return ""
	}

	// 处理指针类型
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.ValueOf(i).IsNil() {
			return ""
		}
		i = reflect.ValueOf(i).Elem().Interface()
	}

	switch v := i.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case complex64:
		return fmt.Sprintf("(%g+%gi)", real(v), imag(v))
	case complex128:
		return fmt.Sprintf("(%g+%gi)", real(v), imag(v))
	case bool:
		return strconv.FormatBool(v)
	//case *string:
	//	if v == nil {
	//		return ""
	//	}
	//	return *v
	//case *int:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.Itoa(*v)
	//case *int8:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatInt(int64(*v), 10)
	//case *int16:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatInt(int64(*v), 10)
	//case *int32:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatInt(int64(*v), 10)
	//case *int64:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatInt(*v, 10)
	//case *uint:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatUint(uint64(*v), 10)
	//case *uint8:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatUint(uint64(*v), 10)
	//case *uint16:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatUint(uint64(*v), 10)
	//case *uint32:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatUint(uint64(*v), 10)
	//case *uint64:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatUint(*v, 10)
	//case *float32:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatFloat(float64(*v), 'f', -1, 32)
	//case *float64:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatFloat(*v, 'f', -1, 64)
	//case *complex64:
	//	if v == nil {
	//		return ""
	//	}
	//	return fmt.Sprintf("(%g+%gi)", real(*v), imag(*v))
	//case *complex128:
	//	if v == nil {
	//		return ""
	//	}
	//	return fmt.Sprintf("(%g+%gi)", real(*v), imag(*v))
	//case *bool:
	//	if v == nil {
	//		return ""
	//	}
	//	return strconv.FormatBool(*v)
	default:
		return ""
	}
}

func toStringReflect(i any) string {
	if i == nil {
		return ""
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Complex64:
		return fmt.Sprintf("(%g+%gi)", real(v.Complex()), imag(v.Complex()))
	case reflect.Complex128:
		return fmt.Sprintf("(%g+%gi)", real(v.Complex()), imag(v.Complex()))
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	default:
		return ""
	}
}
