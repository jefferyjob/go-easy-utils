package jsonUtil

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

var (
	// ErrPoint 不是指针类型
	ErrPoint = errors.New("the argument to Result must be a non-nil pointer")
	// ErrNotMap 不是Map类型
	ErrNotMap = errors.New("cannot parse map, value is not a map")
	// ErrNotSlice 不是Slice类型
	ErrNotSlice = errors.New("cannot parse slice, value is not a slice")
)

// JsonToStruct Parses JSON into a specified structure pointer
// 将JSON解析为指定的结构体指针
func JsonToStruct(jsonData string, result interface{}) error {
	if reflect.ValueOf(result).Kind() != reflect.Pointer || reflect.ValueOf(result).IsNil() {
		return ErrPoint
	}

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

		// 获取json的tag
		jsonTag := getJsonTag(fieldType, fieldName)

		value, ok := data[jsonTag]
		if !ok {
			continue
		}

		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(toStringReflect(value))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val, err := toInt64Reflect(value)
			if err != nil {
				log.Printf("toInt64Reflect err:%s \n", err)
				return err
			}
			fieldValue.SetInt(val)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val, err := toUint64Reflect(value)
			if err != nil {
				log.Printf("toUint64Reflect err:%s \n", err)
				return err
			}
			fieldValue.SetUint(val)
		case reflect.Float32, reflect.Float64:
			val, err := toFloat64Reflect(value)
			if err != nil {
				log.Printf("toFloat64Reflect err:%s \n", err)
				return err
			}
			fieldValue.SetFloat(val)
		case reflect.Bool:
			val := toBoolReflect(value)
			fieldValue.SetBool(val)
		case reflect.Struct:
			if subData, ok := value.(map[string]any); ok {
				subResult := reflect.New(fieldValue.Type())
				err := JsonToStruct(convertToJSONString(subData), subResult.Interface())
				if err != nil {
					return err
				}
				fieldValue.Set(subResult.Elem())
			}
		case reflect.Map:
			if subData, ok := value.(map[string]interface{}); ok {
				subResult := reflect.New(fieldType.Type).Elem()
				if err := parseMap(subResult, subData); err != nil {
					return err
				}
				fieldValue.Set(subResult)
			}
		case reflect.Slice:
			if subData, ok := value.([]interface{}); ok {
				if err := parseSlice(fieldValue, subData); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("unexpected value type for slice: %T", value)
			}
		case reflect.Interface:
			if value == nil {
				fieldValue.Set(reflect.Zero(fieldType.Type))
			} else {
				fieldValue.Set(reflect.ValueOf(value))
			}
		}
	}
	return nil
}

// 从json的tag标签中取出定义字段
func getJsonTag(fieldType reflect.StructField, fieldName string) string {
	jsonTag := fieldType.Tag.Get("json")
	if jsonTag == "" {
		jsonTag = fieldName
	} else {
		if commaIndex := strings.Index(jsonTag, ","); commaIndex != -1 {
			jsonTag = jsonTag[:commaIndex]
		}
	}
	return jsonTag
}

func convertToJSONString(data map[string]interface{}) string {
	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes)
}
