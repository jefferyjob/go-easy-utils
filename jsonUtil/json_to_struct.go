package jsonUtil

import (
	"encoding/json"
	"github.com/jefferyjob/go-easy-utils"
	"reflect"
	"strconv"
	"strings"
)

// JsonToStruct 将 JSON 字符串解析为指定的结构体指针
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

		// 从json的tag标签中取出定义字段
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = fieldName
		} else {
			if commaIndex := strings.Index(jsonTag, ","); commaIndex != -1 {
				jsonTag = jsonTag[:commaIndex]
			}
		}

		value, ok := data[jsonTag]
		if !ok {
			continue
		}

		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(toString(value))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val, err := toInt64(value)
			if err != nil {
				return err
			}
			fieldValue.SetInt(val)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val, err := toUint64(value)
			if err != nil {
				return err
			}
			fieldValue.SetUint(val)
		case reflect.Float32, reflect.Float64:
			val, err := toFloat64(value)
			if err != nil {
				return err
			}
			fieldValue.SetFloat(val)
		case reflect.Bool:
			val := toBool(value)
			fieldValue.SetBool(val)
		case reflect.Struct:
			if subData, ok := value.(map[string]interface{}); ok {
				subResult := reflect.New(fieldValue.Type())
				err := JsonToStruct(convertToJSONString(subData), subResult.Interface())
				if err != nil {
					return err
				}
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
							err := JsonToStruct(convertToJSONString(subDataElem), subResultElem.Interface())
							if err != nil {
								return err
							}
							subElem.Set(subResultElem.Elem())
						}
					} else {
						subElem.Set(reflect.ValueOf(subValue))
					}
				}
				fieldValue.Set(subResult)
			}
			//default:
			//	fieldValue.Set(reflect.ValueOf(value))
		}
	}

	return nil
}

func convertToJSONString(data map[string]interface{}) string {
	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes)
}

func toBool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case string:
		return v != "" && v != "false" && v != "0"
	case int, int8, int16, int32, int64:
		return v != 0
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return v != 0
	case float32, float64:
		return v != 0.0
	case complex64, complex128:
		return v != 0+0i
	case nil:
		return false
	default:
		return false
	}
}

func toString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}

func toInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, nil
	}

	switch value.(type) {
	case float32:
		return int64(value.(float32)), nil
	case float64:
		return int64(value.(float64)), nil
	case string:
		intValue, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return 0, go_easy_utils.ErrSyntax
		}
		return intValue, nil
	case int:
		return int64(value.(int)), nil
	case int8:
		return int64(value.(int8)), nil
	case int16:
		return int64(value.(int16)), nil
	case int32:
		return int64(value.(int32)), nil
	case int64:
		return value.(int64), nil
	case interface{}:
		return 0, nil
	default:
		return 0, go_easy_utils.ErrType
	}
}

func toUint64(value interface{}) (uint64, error) {
	if value == nil {
		return 0, nil
	}

	switch value.(type) {
	case float32:
		v := value.(float32)
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case float64:
		v := value.(float64)
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case string:
		intValue, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return 0, go_easy_utils.ErrSyntax
		}
		return intValue, nil
	case uint:
		return uint64(value.(uint)), nil
	case uint8:
		return uint64(value.(uint8)), nil
	case uint16:
		return uint64(value.(uint16)), nil
	case uint32:
		return uint64(value.(uint32)), nil
	case uint64:
		return value.(uint64), nil
	case int:
		v := value.(int)
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case int8:
		v := value.(int8)
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case int16:
		v := value.(int16)
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case int32:
		v := value.(int32)
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case int64:
		v := value.(int64)
		if v < 0 {
			return 0, nil
		}
		return uint64(v), nil
	case interface{}:
		return 0, nil
	default:
		return 0, go_easy_utils.ErrType
	}
}

func toFloat64(value interface{}) (float64, error) {
	if value == nil {
		return 0, nil
	}

	switch value.(type) {
	case float32:
		return float64(value.(float32)), nil
	case float64:
		return value.(float64), nil
	case string:
		floatValue, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			return 0, go_easy_utils.ErrSyntax
		}
		return floatValue, nil
	default:
		return 0, go_easy_utils.ErrType
	}
}
