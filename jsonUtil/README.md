# jsonUtil

根据结构体的字段类型和标签来自动选择将 JSON 值转换为相应的类型。

支持的字段类型包括 string、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、bool、float32 和 float64。

支持的标签有 "json"、"jsonb" 和 "mapstructure"。
- "json" 和 "jsonb" 标签指示解析 JSON 时使用的键名。
- "mapstructure" 标签指示字段名的映射关系。

如果 JSON 中的某些键在结构体中没有对应的字段，则它们将被忽略。
如果 JSON 中的某些键的类型与结构体中的字段类型不匹配，则会引发解析错误。

参数 jsonData 是要解析的 JSON 字符串。
参数 result 是指向要填充 JSON 值的结构体指针。

如果解析成功，则返回 nil。如果解析失败，则返回解析错误。

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/jsonUtil
```

## ErrorCode


## Functions

```go
// JsonToStruct 将 JSON 字符串解析为指定的结构体指针
func JsonToStruct(jsonData string, result interface{}) error
```