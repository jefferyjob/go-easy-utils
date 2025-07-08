package jsonx

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func parseSlice(fieldValue reflect.Value, subData []any) error {
	if fieldValue.Kind() != reflect.Slice {
		return ErrNotSlice
	}

	// 获取切片元素类型
	elemType := fieldValue.Type().Elem()

	// 创建一个与目标字段相同类型的新切片
	slice := reflect.MakeSlice(fieldValue.Type(), 0, len(subData))

	// 迭代 subData 切片的元素
	for _, item := range subData {
		// 创建 slice 元素类型的新元素
		elem := reflect.New(elemType).Elem()

		// 如果元素类型是结构体，则递归调用 JsonToStruct
		if elem.Kind() == reflect.Struct {
			// 将项目转换为 JSON 字符串
			jsonData, err := json.Marshal(item)
			if err != nil {
				return err
			}

			// 调用 ToStruct 将 JSON 字符串解析为 struct 元素
			err = ToStruct(string(jsonData), elem.Addr().Interface())
			if err != nil {
				return err
			}
		} else if elemType.Kind() == reflect.Interface {
			if item == nil {
				elem.Set(reflect.Zero(elemType))
			} else {
				elem.Set(reflect.ValueOf(item))
			}
		} else {
			// 否则，将项目转换为适当的类型并设置元素值
			switch elem.Kind() {
			case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
				reflect.Float32, reflect.Float64, reflect.Bool:
				err := parsePrimitiveValue(elem, item)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unsupported slice element type: %s", elemType.String())
			}
		}

		slice = reflect.Append(slice, elem) // 将元素附加到切片
	}

	fieldValue.Set(slice) // 将切片值设置为目标字段

	return nil
}
