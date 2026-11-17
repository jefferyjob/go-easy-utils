package jsonx

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// normalizeNumber 将 json.Number 归一化为后续可解析的标量：
// 整数字面量转为 string 以精确解析（避免 float64 丢精度）；
// 含小数或指数的字面量转为 float64，保持与浮点语义一致。
func normalizeNumber(value any) (any, error) {
	n, ok := value.(json.Number)
	if !ok {
		return value, nil
	}
	if strings.ContainsAny(string(n), ".eE") {
		f, err := n.Float64()
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	return string(n), nil
}

func setBool(field reflect.Value, value any, path string) error {
	switch v := value.(type) {
	case bool:
		field.SetBool(v)
	case string:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return fmt.Errorf("%s: cannot convert %v to bool, err: %v", path, v, err)
		}
		field.SetBool(b)
	}
	return nil
}

func setInt(field reflect.Value, value any, path string) error {
	value, err := normalizeNumber(value)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	// 按字段实际位宽解析，超出范围时返回错误而不是静默截断
	bits := field.Type().Bits()
	switch v := value.(type) {
	case float64:
		limit := math.Ldexp(1, bits-1)
		if v < -limit || v >= limit {
			return fmt.Errorf("%s: %v 超出 %s 范围: %w", path, v, field.Type(), ErrValOut)
		}
		field.SetInt(int64(v))
	case string:
		i, err := strconv.ParseInt(v, 10, bits)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				return fmt.Errorf("%s: %q 超出 %s 范围: %w", path, v, field.Type(), ErrValOut)
			}
			return fmt.Errorf("%s: %w", path, err)
		}
		field.SetInt(i)
	default:
		return fmt.Errorf("%s: 无法转换 %v 为 int", path, value)
	}
	return nil
}

func setUint(field reflect.Value, value any, path string) error {
	value, err := normalizeNumber(value)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	// 按字段实际位宽解析，负数或超出范围时返回错误而不是静默绕回
	bits := field.Type().Bits()
	switch v := value.(type) {
	case float64:
		if v < 0 || v >= math.Ldexp(1, bits) {
			return fmt.Errorf("%s: %v 超出 %s 范围: %w", path, v, field.Type(), ErrValOut)
		}
		field.SetUint(uint64(v))
	case string:
		u, err := strconv.ParseUint(v, 10, bits)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				return fmt.Errorf("%s: %q 超出 %s 范围: %w", path, v, field.Type(), ErrValOut)
			}
			return fmt.Errorf("%s: %w", path, err)
		}
		field.SetUint(u)
	default:
		return fmt.Errorf("%s: 无法转换 %v 为 uint", path, value)
	}
	return nil
}

func setFloat(field reflect.Value, value any, path string) error {
	// json.Number 直接按字面量解析，按字段位宽单次舍入，避免双重舍入丢精度
	if n, ok := value.(json.Number); ok {
		value = string(n)
	}
	// 按字段实际位宽解析，超出范围时返回错误而不是静默变 Inf 或丢失精度
	bits := field.Type().Bits()
	switch v := value.(type) {
	case float64:
		// 窄化为 float32 时若溢出会得到 Inf，需显式检查
		if bits < 64 && math.IsInf(float64(float32(v)), 0) && !math.IsInf(v, 0) {
			return fmt.Errorf("%s: %v 超出 %s 范围: %w", path, v, field.Type(), ErrValOut)
		}
		field.SetFloat(v)
	case string:
		f, err := strconv.ParseFloat(v, bits)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				return fmt.Errorf("%s: %q 超出 %s 范围: %w", path, v, field.Type(), ErrValOut)
			}
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
		return fmt.Errorf("%s: 不是数组类型", path)
	}

	elemType := field.Type().Elem()
	slice := reflect.MakeSlice(field.Type(), 0, len(valSlice))

	for i, item := range valSlice {
		elem := reflect.New(elemType).Elem()
		// 走 setFieldValue 以支持指针元素（如 []*int）及 null 元素
		if err := setFieldValue(elem, elemType, item, fmt.Sprintf("%s[%d]", path, i)); err != nil {
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

		// 递归处理 value，走 setFieldValue 以支持指针值（如 map[string]*int）及 null 值
		elem := reflect.New(elemType).Elem()
		if err := setFieldValue(elem, elemType, rawVal, fmt.Sprintf("%s[%s]", path, rawKey)); err != nil {
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
