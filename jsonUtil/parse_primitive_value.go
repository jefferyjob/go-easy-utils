package jsonUtil

import (
	"errors"
	"fmt"
	"reflect"
)

func parsePrimitiveValue(fieldVal reflect.Value, v any) error {
	switch fieldVal.Kind() {
	case reflect.String:
		str, ok := v.(string)
		if !ok {
			return errors.New("failed to parse string")
		}
		fieldVal.SetString(str)
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
		b, ok := v.(bool)
		if !ok {
			return errors.New("failed to parse bool")
		}
		fieldVal.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind: %s", fieldVal.Kind())
	}
	return nil
}

//case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
//reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
//reflect.Float32, reflect.Float64, reflect.Bool:
//if err := parsePrimitiveValue(fieldVal, v); err != nil {
//return err
//}
