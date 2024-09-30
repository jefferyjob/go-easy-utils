# jsonUtil

[English](README.md) | 简体中文

## 介绍

JsonToStruct 将 JSON 数据转换为 Go 结构。

**参数：**
- jsonData：包含 JSON 数据的字符串。
- val：指向要填充的结构变量的指针。

**返回：**
- error：如果转换失败或发生错误，则返回相应的错误。如果成功，则返回 nil。

**功能：**
- 检查 val 参数是否为非零指针类型，如果不是，则返回 ErrPoint。
- 将 jsonData 解析为名为 data 的 map[string]any 变量。
- 使用反射检索 val 指向的结构的值和类型。
- 遍历结构的字段：
- 检索字段的类型、名称和值。
- 获取字段的 JSON 标签。
- 如果数据中存在与 JSON 标签对应的键值对，则进行相应的处理：
- 如果字段是原始类型（string、integer、float、boolean），则解析为对应类型的值。
- 如果字段是结构体类型，则递归调用 JsonToStruct 函数将子结构体转换为 JSON。
- 如果字段是映射类型，则使用 parseMap 函数将子映射转换为 JSON。
- 如果字段是切片类型，则使用 parseSlice 函数将子切片转换为 JSON。
- 如果字段是接口类型，则将值设置为 nil 或对应值。

## 安装

```bash
go get -u github.com/jefferyjob/go-easy-utils/v2/jsonUtil
```

## 导入

```go
import (
	"github.com/jefferyjob/go-easy-utils/jsonUtil"
)
```

## 方法

```go
func JsonToStruct(jsonData string, val any) error
```

## Demo

```go
// name 使用了两个 json 标签
// age 定义为 int，json 的值是 string
// is_use 定义为 bool，json 的值是 int
func TestDemo1(t *testing.T) {
    jsonData := `{
        "name": "make",
        "age": "22",
        "is_use": "1"
    }`
    
    var people struct {
        Name  string `json:"name,omitempty"`
        Age   int    `json:"age"`
        IsUse bool   `json:"is_use"`
    }
    
    if err := JsonToStruct(jsonData, &people); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%+v \n", people)
    // return
    // {Name:make Age:22 IsUse:true}
}
```

```go
// 结构嵌套和切片嵌套处理
func TestJsonToStructDemo2(t *testing.T) {
	type Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	}
	type Person struct {
		Name      string   `json:"name"`
		Age       int      `json:"age"`
		Address   Address  `json:"address"`
		Interests []string `json:"interests"`
	}

	jsonData := `{
            "name": "Bob",
            "age": "25",
            "address": {
                "city": "Shanghai",
                "country": "China"
            },
            "interests": ["reading", "swimming"]
        }`

	var person Person
	err := JsonToStruct(jsonData, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v \n", person)
	// {Name:Bob Age:25 Address:{City:Shanghai Country:China} Interests:[reading swimming]}
}
```
