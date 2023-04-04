package sliceUtil

import (
	"reflect"
)

// ColumnSlice 获取slice中某个单一列的值
func ColumnSlice[T any](slice []T, column string) []any {
	// 获取切片的反射值
	sliceValue := reflect.ValueOf(slice)

	// 获取切片的长度和元素类型
	length := sliceValue.Len()
	elementType := sliceValue.Type().Elem()

	// 获取指定字段的反射值
	fieldIndex := -1
	for i := 0; i < elementType.NumField(); i++ {
		field := elementType.Field(i)
		if field.Name == column {
			fieldIndex = i
			break
		}
	}
	if fieldIndex == -1 {
		return []any{}
	}

	// 构造返回值切片
	result := make([]any, length)

	// 获取每个元素指定字段的值，并添加到返回值切片中
	for i := 0; i < length; i++ {
		elementValue := sliceValue.Index(i)
		fieldValue := elementValue.Field(fieldIndex)
		result[i] = fieldValue.Interface()
	}

	return result
}
