package jsonUtil

import (
	"fmt"
	"reflect"
)

// 根据用户定义的结构体字段类型，将输入值 v 转换为与该字段匹配的类型，并赋值给该结构体字段
//
// 参数说明:
//   - fieldVal: 目标结构体的字段（通过反射机制传入），其类型决定了如何转换和赋值
//   - v: 需要转换的原始数据
//
// 该方法是对基础类型转换功能（如 `to_primitive.go` 中的方法）的进一步封装，用于在处理结构体时更方便地进行数据类型的自动匹配和转换。
// 原理：
//   - 通过反射获取结构体字段的类型（`fieldVal.Kind()`），并根据不同的类型调用相应的转换方法（如 `toStringReflect`、`toInt64Reflect` 等）。
//   - 如果转换成功，将转换后的值赋值给结构体的字段。
//   - 支持的类型包括：string、int（及其变种）、uint（及其变种）、float（float32、float64）、bool。
//   - 对于不支持的类型，将返回错误。
//
// 使用场景：
//
//	适用于在运行时根据结构体的字段类型，对输入数据进行动态类型转换并赋值的场景。
func parsePrimitiveValue(fieldVal reflect.Value, v any) error {
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
