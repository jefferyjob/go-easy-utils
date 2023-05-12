package jsonUtil

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func parseSlice(fieldValue reflect.Value, subData []interface{}) error {
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

			// 调用 JsonToStruct 将 JSON 字符串解析为 struct 元素
			err = JsonToStruct(string(jsonData), elem.Addr().Interface())
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
			case reflect.String:
				elem.SetString(toStringReflect(item))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				val, err := toInt64Reflect(item)
				if err != nil {
					log.Printf("toInt64Reflect err:%s \n", err)
					return err
				}
				elem.SetInt(val)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				val, err := toUint64Reflect(item)
				if err != nil {
					log.Printf("toUint64Reflect err:%s \n", err)
					return err
				}
				elem.SetUint(val)
			case reflect.Float32, reflect.Float64:
				val, err := toFloat64Reflect(item)
				if err != nil {
					log.Printf("toFloat64Reflect err:%s \n", err)
					return err
				}
				elem.SetFloat(val)
			case reflect.Bool:
				val := toBoolReflect(item)
				elem.SetBool(val)
			default:
				return fmt.Errorf("unsupported slice element type: %s", elemType.String())
			}
		}

		// 将元素附加到切片
		slice = reflect.Append(slice, elem)
	}

	// 将切片值设置为目标字段
	fieldValue.Set(slice)

	return nil
}
