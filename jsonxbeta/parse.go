package jsonxbeta

import (
	"fmt"
	"reflect"
	"strconv"
)

func setBool(field reflect.Value, value any, path string) error {
	switch v := value.(type) {
	case bool:
		field.SetBool(v)
	case string:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return fmt.Errorf("%s: cannot convert %v to bool", path, v)
		}
		field.SetBool(b)
	}
	return nil
}

func setInt(field reflect.Value, value any, path string) error {
	switch v := value.(type) {
	case float64:
		field.SetInt(int64(v))
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		field.SetInt(i)
	default:
		return fmt.Errorf("%s: 无法转换 %v 为 int", path, value)
	}
	return nil
}

func setUint(field reflect.Value, value any, path string) error {
	switch v := value.(type) {
	case float64:
		field.SetUint(uint64(v))
	case string:
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		field.SetUint(u)
	default:
		return fmt.Errorf("%s: 无法转换 %v 为 uint", path, value)
	}
	return nil
}

func setFloat(field reflect.Value, value any, path string) error {
	switch v := value.(type) {
	case float64:
		field.SetFloat(v)
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		field.SetFloat(f)
	default:
		return fmt.Errorf("%s: 无法转换 %v 为 float", path, value)
	}
	return nil
}

func setSlice(field reflect.Value, value any, path string) error {
	valSlice, ok := value.([]any)
	if !ok {
		return fmt.Errorf("不是数组类型")
	}

	elemType := field.Type().Elem()
	slice := reflect.MakeSlice(field.Type(), 0, len(valSlice))

	for i, item := range valSlice {
		elem := reflect.New(elemType).Elem()
		if err := assignValue(elem, item, fmt.Sprintf("%s[%d]", path, i)); err != nil {
			return err
		}
		slice = reflect.Append(slice, elem)
	}
	field.Set(slice)
	return nil
}

// func setMap(field reflect.Value, value any, path string) error {
// 	if field.Type().Key().Kind() != reflect.String {
// 		return fmt.Errorf("map 仅支持 string 类型的 key")
// 	}
//
// 	valMap, ok := value.(map[string]any)
// 	if !ok {
// 		return fmt.Errorf("%s: 不是 map 类型", path)
// 	}
//
// 	mapValue := reflect.MakeMap(field.Type())
// 	elemType := field.Type().Elem()
//
// 	for k, v := range valMap {
// 		elem := reflect.New(elemType).Elem()
// 		if err := assignValue(elem, v, fmt.Sprintf("%s[%s]", path, k)); err != nil {
// 			return err
// 		}
// 		mapValue.SetMapIndex(reflect.ValueOf(k), elem)
// 	}
// 	field.Set(mapValue)
// 	return nil
// }

// setMap 支持任意基础类型作为 key
func setMap(field reflect.Value, value any, path string) error {
	// JSON decode 后对象总是 map[string]any
	rawMap, ok := value.(map[string]any)
	if !ok {
		return fmt.Errorf("%s: 不是 map 类型", path)
	}

	fieldType := field.Type()
	keyType := fieldType.Key()
	elemType := fieldType.Elem()

	// 新建目标类型的 map
	result := reflect.MakeMapWithSize(fieldType, len(rawMap))

	for rawKey, rawVal := range rawMap {
		// 把 string key 转成目标类型
		keyVal, err := parseMapKey(rawKey, keyType, path)
		if err != nil {
			return err
		}

		// 递归处理 value
		elem := reflect.New(elemType).Elem()
		if err := assignValue(elem, rawVal, fmt.Sprintf("%s[%s]", path, rawKey)); err != nil {
			return err
		}

		result.SetMapIndex(keyVal, elem)
	}

	field.Set(result)
	return nil
}

// parseMapKey 将 JSON 的 string key 转换为目标 key 类型的 reflect.Value
func parseMapKey(keyStr string, keyType reflect.Type, path string) (reflect.Value, error) {
	switch keyType.Kind() {
	case reflect.String:
		return reflect.ValueOf(keyStr).Convert(keyType), nil
	case reflect.Bool:
		b, err := strconv.ParseBool(keyStr)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("%s: 无法将 %q 转为 bool", path, keyStr)
		}
		return reflect.ValueOf(b).Convert(keyType), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		bitSize := keyType.Bits()
		i, err := strconv.ParseInt(keyStr, 10, bitSize)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("%s: 无法将 %q 转为 %s", path, keyStr, keyType.Kind())
		}
		return reflect.ValueOf(i).Convert(keyType), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		bitSize := keyType.Bits()
		u, err := strconv.ParseUint(keyStr, 10, bitSize)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("%s: 无法将 %q 转为 %s", path, keyStr, keyType.Kind())
		}
		return reflect.ValueOf(u).Convert(keyType), nil
	case reflect.Float32, reflect.Float64:
		bitSize := keyType.Bits()
		f, err := strconv.ParseFloat(keyStr, bitSize)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("%s: 无法将 %q 转为 %s", path, keyStr, keyType.Kind())
		}
		return reflect.ValueOf(f).Convert(keyType), nil
	default:
		return reflect.Value{}, fmt.Errorf("%s: 不支持的 map key 类型 %s", path, keyType.Kind())
	}
}
