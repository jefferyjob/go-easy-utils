package jsonUtil

import (
	"errors"
	"fmt"
	"reflect"
)

func parseMap(value reflect.Value, data map[string]any) error {
	if value.Kind() != reflect.Map {
		return ErrNotMap
	}

	valueType := value.Type().Elem()
	newMap := reflect.MakeMapWithSize(value.Type(), len(data))

	for key, val := range data {
		newKey := reflect.ValueOf(key)

		if valueType.Kind() == reflect.Interface {
			newMap.SetMapIndex(newKey, reflect.ValueOf(val))
		} else if valueType.Kind() == reflect.Map {
			newElem := reflect.New(valueType).Elem()
			if err := parseMap(newElem, val.(map[string]any)); err != nil {
				return err
			}
			newMap.SetMapIndex(newKey, newElem)
		} else {
			newVal := reflect.New(valueType).Elem()
			if err := parseValue(newVal, val); err != nil {
				return err
			}
			newMap.SetMapIndex(newKey, newVal)
		}
	}

	value.Set(newMap)
	return nil
}

func parseValue(fieldVal reflect.Value, v any) error {
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
	case reflect.Struct:
		if subData, ok := v.(map[string]any); ok {
			if err := JsonToStruct(convertToJSONString(subData), fieldVal.Addr().Interface()); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unexpected value type for struct: %T", v)
		}
	case reflect.Slice:
		if arr, ok := v.([]any); ok {
			elemType := fieldVal.Type().Elem()
			sliceVal := reflect.MakeSlice(fieldVal.Type(), len(arr), len(arr))
			for i, elem := range arr {
				elemVal := reflect.New(elemType).Elem()
				if err := parseValue(elemVal, elem); err != nil {
					return err
				}
				sliceVal.Index(i).Set(elemVal)
			}
			fieldVal.Set(sliceVal)
		} else {
			return fmt.Errorf("unexpected value type for slice: %T", v)
		}
	case reflect.Map:
		if mapData, ok := v.(map[string]any); ok {
			mapType := fieldVal.Type()
			if fieldVal.IsNil() {
				fieldVal.Set(reflect.MakeMap(mapType))
			}
			elemType := mapType.Elem()
			mapVal := fieldVal
			for k, v := range mapData {
				keyVal := reflect.New(mapType.Key()).Elem()
				if err := parseValue(keyVal, k); err != nil {
					return err
				}
				elemVal := reflect.New(elemType).Elem()
				if err := parseValue(elemVal, v); err != nil {
					return err
				}
				mapVal.SetMapIndex(keyVal, elemVal)
			}
		} else {
			return fmt.Errorf("unexpected value type for map: %T", v)
		}
	default:
		return fmt.Errorf("unsupported kind: %s", fieldVal.Kind())
	}
	return nil
}
