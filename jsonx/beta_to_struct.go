package jsonx

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ToStructBeta JSON转结构体
func ToStructBeta(jsonStr string, target any) error {
	// 将 JSON 解成 map[string]any
	var raw map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &raw); err != nil {
		return err
	}

	// 判断必须是指针
	rv := reflect.ValueOf(target)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return ErrPoint
	}

	// 递归赋值到结构体中
	return fill(raw, reflect.ValueOf(target).Elem(), "")
}

// fill 递归填充结构体字段
func fill(data map[string]any, target reflect.Value, path string) error {
	targetType := target.Type()

	for i := 0; i < target.NumField(); i++ {
		field := target.Field(i)
		// 若字段不可写，则跳过
		if !field.CanSet() {
			continue
		}

		structField := targetType.Field(i)
		jsonTag := structField.Tag.Get("json")

		// 如果 jsonTag 为空或为 "-"，则使用字段名
		// if jsonTag == "" || jsonTag == "-" {
		// 	jsonTag = structField.Name
		// }

		// 取对应的值
		value, ok := data[jsonTag]
		if !ok {
			continue
		}

		// 构造字段的完整路径，用于错误提示追踪。
		// 例如当前字段 jsonTag 为 "city"，上一级 path 为 "address"
		// 那么 newPath 就是 "address.city"，表示这是嵌套结构 address 里的 city 字段。
		// 如果是顶层字段（如 "name"），path 为 ""，则直接使用 jsonTag 作为路径。
		newPath := jsonTag
		if path != "" {
			newPath = path + "." + jsonTag
		}

		err := setFieldValue(field, structField.Type, value, newPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// setFieldValue 根据字段类型处理普通值或指针值
func setFieldValue(field reflect.Value, fieldType reflect.Type, value any, path string) error {
	// 如果是指针类型，则初始化并递归赋值其 Elem()
	if fieldType.Kind() == reflect.Ptr {
		// 创建指针指向的类型实例
		elemType := fieldType.Elem()
		elemValue := reflect.New(elemType).Elem()
		// 递归设置值
		if err := setFieldValue(elemValue, elemType, value, path); err != nil {
			return err
		}
		// 设置指针值
		field.Set(elemValue.Addr())
		return nil
	}

	// 非指针类型直接处理
	return assignValue(field, value, path)
}

// assignValue 实际执行字段值赋值的逻辑
func assignValue(field reflect.Value, value any, path string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(fmt.Sprintf("%v", value))
	case reflect.Bool:
		return setBool(field, value, path)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setInt(field, value, path)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUint(field, value, path)
	case reflect.Float32, reflect.Float64:
		return setFloat(field, value, path)
	case reflect.Slice:
		return setSlice(field, value, path)
	case reflect.Map:
		return setMap(field, value, path)
	case reflect.Struct:
		subMap, ok := value.(map[string]any)
		if !ok {
			// 可能是 json.RawMessage 类型，要再解一层
			bs, err := json.Marshal(value)
			if err != nil {
				return fmt.Errorf("%s: %w", path, err)
			}
			var m map[string]any
			if err := json.Unmarshal(bs, &m); err != nil {
				return fmt.Errorf("%s: %w", path, err)
			}
			subMap = m
		}
		return fill(subMap, field, path)
	case reflect.Interface:
		if value == nil {
			field.Set(reflect.Zero(field.Type()))
		} else {
			// 如果是 interface 类型，直接设置值
			field.Set(reflect.ValueOf(value))
		}
	default:
		return fmt.Errorf("%s: 不支持类型 %s", path, field.Kind())
	}

	return nil
}
