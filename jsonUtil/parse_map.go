package jsonUtil

import (
	"fmt"
	"reflect"
)

// 将 JSON 数据中的 map 类型值解析并填充到目标结构的 map 字段中
//
// 参数
//   - value：用户需要解析的目标结构体
//   - data：解析传入的json字符串得到的 `map[string]any` 对应
//
// 检查传入的值是否是 map 类型，如果不是，则返回错误。
//
//	创建一个新的 map，其类型与目标结构字段的类型相同。
//	遍历 JSON 数据中的 map（即 data），对每个键值对进行处理：
//	如果 map 的值类型是接口类型，直接设置对应的值。
//	如果 map 的值类型是另一个 map，递归调用 parseMap 进行处理。
//	对于其他值类型，调用 parseValue 来解析并设置值。
//
// 最后，将处理好的 map 设置到目标结构的对应字段中。
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

// 将 JSON 中的值解析为 Go 中的基本类型或复杂类型
//
// 根据字段的 Kind（数据类型）进行分类处理：
//
//	如果字段是基本类型（如 string、int、float、bool 等），调用 parsePrimitiveValue 进行基础类型的转换。
//	如果字段是结构体，则递归调用 JsonToStruct 来处理嵌套的结构体。
//	如果字段是切片，则创建新的切片，并递归解析每个元素。
//	如果字段是 map，则调用 parseMap 来处理 map 类型。
//	如果字段是 interface{} 类型，根据实际值设置。
//
// 针对不同类型的数据做相应的处理和转换，然后将值设置到目标字段。
func parseValue(fieldVal reflect.Value, item any) error {
	switch fieldVal.Kind() {
	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool:
		if err := parsePrimitiveValue(fieldVal, item); err != nil {
			return err
		}
	case reflect.Struct:
		if subData, ok := item.(map[string]any); ok {
			if err := JsonToStruct(convertToJSONString(subData), fieldVal.Addr().Interface()); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unexpected value type for struct: %T", item)
		}
	case reflect.Slice:
		if arr, ok := item.([]any); ok {
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
			return fmt.Errorf("unexpected value type for slice: %T", item)
		}
	case reflect.Map:
		if mapData, ok := item.(map[string]any); ok {
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
			return fmt.Errorf("unexpected value type for map: %T", item)
		}
	case reflect.Interface:
		if item == nil {
			fieldVal.Set(reflect.Zero(fieldVal.Type()))
		} else {
			fieldVal.Set(reflect.ValueOf(item))
		}
	default:
		return fmt.Errorf("unsupported kind: %s", fieldVal.Kind())
	}
	return nil
}
