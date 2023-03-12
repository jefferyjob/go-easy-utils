// Package jsonUtil Json数据处理包
// 可用于json赋值结构体，json数据转义
package jsonUtil

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// JsonToStruct 将 JSON 字符串解析为指定的结构体指针
// 根据结构体的字段类型和标签来自动选择将 JSON 值转换为相应的类型。
//
// 支持的字段类型包括 string、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、bool、float32 和 float64。
//
// 支持的标签有 "json"、"jsonb" 和 "mapstructure"。
// - "json" 和 "jsonb" 标签指示解析 JSON 时使用的键名。
// - "mapstructure" 标签指示字段名的映射关系。
//
// 如果 JSON 中的某些键在结构体中没有对应的字段，则它们将被忽略。
// 如果 JSON 中的某些键的类型与结构体中的字段类型不匹配，则会引发解析错误。
//
// 参数 jsonData 是要解析的 JSON 字符串。
// 参数 result 是指向要填充 JSON 值的结构体指针。
//
// 如果解析成功，则返回 nil。如果解析失败，则返回解析错误。
func JsonToStruct(jsonData string, result interface{}) error {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return err
	}

	resultValue := reflect.ValueOf(result).Elem()
	resultType := resultValue.Type()

	for i := 0; i < resultType.NumField(); i++ {
		fieldType := resultType.Field(i)
		fieldName := fieldType.Name
		fieldValue := resultValue.FieldByName(fieldName)

		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = fieldName
		}

		value, ok := data[jsonTag]
		if !ok {
			continue
		}

		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(value.(string))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			switch value.(type) {
			case float64:
				fieldValue.SetInt(int64(value.(float64)))
			case string:
				if intValue, err := strconv.ParseInt(value.(string), 10, 64); err == nil {
					fieldValue.SetInt(intValue)
				}
			default:
				fieldValue.SetInt(int64(value.(int)))
			}
		case reflect.Float32, reflect.Float64:
			switch value.(type) {
			case float64:
				fieldValue.SetFloat(value.(float64))
			case string:
				if floatValue, err := strconv.ParseFloat(value.(string), 64); err == nil {
					fieldValue.SetFloat(floatValue)
				}
			}
		case reflect.Struct:
			if subData, ok := value.(map[string]interface{}); ok {
				subResult := reflect.New(fieldValue.Type())
				JsonToStruct(convertToJSONString(subData), subResult.Interface())
				fieldValue.Set(subResult.Elem())
			}
		case reflect.Slice:
			if subData, ok := value.([]interface{}); ok {
				subResult := reflect.MakeSlice(fieldValue.Type(), len(subData), len(subData))
				for j := 0; j < len(subData); j++ {
					subValue := subData[j]
					subElem := subResult.Index(j)

					if subElem.Kind() == reflect.Struct {
						if subDataElem, ok := subValue.(map[string]interface{}); ok {
							subResultElem := reflect.New(subElem.Type())
							JsonToStruct(convertToJSONString(subDataElem), subResultElem.Interface())
							subElem.Set(subResultElem.Elem())
						}
					} else {
						subElem.Set(reflect.ValueOf(subValue))
					}
				}
				fieldValue.Set(subResult)
			}
		default:
			fieldValue.Set(reflect.ValueOf(value))
		}
	}

	return nil
}

func convertToJSONString(data map[string]interface{}) string {
	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes)
}
