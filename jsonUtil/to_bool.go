package jsonUtil

import (
	"fmt"
	"reflect"
)

func toBool(i interface{}) bool {
	if i == nil {
		return false
	}

	// 处理指针类型
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.ValueOf(i).IsNil() {
			return false
		}
		i = reflect.ValueOf(i).Elem().Interface()
	}

	switch v := i.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case int8:
		return v != 0
	case int16:
		return v != 0
	case int32:
		return v != 0
	case int64:
		return v != 0
	case uint:
		return v != 0
	case uint8:
		return v != 0
	case uint16:
		return v != 0
	case uint32:
		return v != 0
	case uint64:
		return v != 0
	case uintptr:
		return v != 0
	case float32:
		return v != 0
	case float64:
		return v != 0
	case complex64:
		return v != 0
	case complex128:
		return v != 0
	case string:
		if v == "true" {
			return true
		} else if v == "false" {
			return false
		}
		return v != ""
	case fmt.Stringer:
		return v.String() != ""
	case interface{ IsNil() bool }:
		return !v.IsNil()
	default:
		return false
	}
}

func toBoolReflect(i interface{}) bool {
	if i == nil {
		return false
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return false
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() != 0
	case reflect.String:
		val := v.String()
		if val == "true" {
			return true
		} else if val == "false" {
			return false
		}
		return v.String() != ""
	//case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice:
	//	return !v.IsNil()
	default:
		return false
	}
}
