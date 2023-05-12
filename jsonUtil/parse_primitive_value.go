package jsonUtil

import (
	"fmt"
	"reflect"
)

// 原始数据类型转换支持
// 将数据转为Struct指定的类型
func parsePrimitiveValue(fieldVal reflect.Value, v interface{}) error {
	switch fieldVal.Kind() {
	case reflect.String:
		fieldVal.SetString(toStringReflect(v))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := toInt64Reflect(v)
		if err != nil {
			return err
		}
		fieldVal.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := toUint64Reflect(v)
		if err != nil {
			return err
		}
		fieldVal.SetUint(uint64(n))
	case reflect.Float32, reflect.Float64:
		n, err := toFloat64Reflect(v)
		if err != nil {
			return err
		}
		fieldVal.SetFloat(n)
	case reflect.Bool:
		fieldVal.SetBool(toBoolReflect(v))
	default:
		return fmt.Errorf("unsupported kind: %s", fieldVal.Kind())
	}
	return nil
}
