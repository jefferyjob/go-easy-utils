package jsonx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// ToStruct JSON转结构体
func ToStruct(jsonStr string, target any) error {
	// UseNumber 让数字保留为 json.Number 字面量，避免超过 2^53 的整数精度丢失
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	dec.UseNumber()

	// 将 JSON 解成 map[string]any
	var raw map[string]any
	if err := dec.Decode(&raw); err != nil {
		return err
	}

	// 与 json.Unmarshal 保持一致：拒绝首个 JSON 值之后的尾随数据
	if _, err := dec.Token(); err != io.EOF {
		if err != nil {
			return err
		}
		return fmt.Errorf("unexpected data after top-level JSON value: %w", ErrSyntax)
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

		// json:"-" 表示显式忽略该字段
		if jsonTag == "-" {
			continue
		}

		// tag 可携带选项（如 "name,omitempty"），匹配时只取逗号前的名称部分
		if idx := strings.Index(jsonTag, ","); idx >= 0 {
			jsonTag = jsonTag[:idx]
		}

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
	// JSON null 保持字段原值：指针字段保持 nil，非指针字段保持零值，与 encoding/json 语义一致
	if value == nil {
		return nil
	}

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
			// 与 ToStruct 一致使用 UseNumber，保持数字精度
			dec := json.NewDecoder(bytes.NewReader(bs))
			dec.UseNumber()
			if err := dec.Decode(&m); err != nil {
				return fmt.Errorf("%s: %w", path, err)
			}
			subMap = m
		}
		return fill(subMap, field, path)
	case reflect.Interface:
		// 值的具体类型必须可赋值给目标接口，否则返回错误而不是让 reflect.Set panic
		rv := reflect.ValueOf(value)
		if rv.Type().AssignableTo(field.Type()) {
			field.Set(rv)
		} else {
			return errors.Join(ErrUnsupported, fmt.Errorf("%s: 类型 %s 无法赋值给接口类型 %s", path, rv.Type(), field.Type()))
		}
	default:
		return errors.Join(ErrUnsupported, fmt.Errorf("%s: 不支持类型 %s", path, field.Kind()))
	}

	return nil
}
