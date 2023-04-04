# jsonUtil

## Introduce

<details>
<summary>English</summary>
JsonToStruct is a method that parses a JSON string into a struct. This method accepts two parameters:  

- jsonData: the JSON string to be parsed  
- result: a pointer to an instance of a struct used to store the parsed data  

The method uses reflection to parse the struct and extracts the corresponding values from the JSON based on the fields defined in the struct and their corresponding JSON tags. If a field in the struct is a nested struct, it recursively parses the nested struct and stores the result in the parent struct.  

This method can handle basic data types such as strings, integers, floating-point numbers, booleans, as well as nested structs and slices. If a value in the JSON string cannot be converted to the target type, the method will return an error.  
</details>

<details>
<summary>简体中文</summary>
JsonToStruct 是一个将JSON字符串解析为结构体的方法。这个方法接受两个参数:  

- jsonData：待解析的JSON字符串  
- result：用于存储解析后数据的结构体实例的指针  

该方法使用了反射机制来解析结构体，并根据结构体中定义的字段和对应的json标签从JSON中提取对应的值。如果结构体的字段是一个嵌套的结构体，它将递归解析嵌套的结构体，并将结果存储在父结构体中。 

该方法可以处理基本数据类型，如字符串、整数、浮点数、布尔值以及嵌套的结构体和切片。如果JSON字符串中的值无法转换为目标类型，该方法将返回一个错误。  
</details>


## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/jsonUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/jsonUtil"
)
```

## Functions

```go
// JsonToStruct Parses JSON into a specified structure pointer
// 将JSON解析为指定的结构体指针
func JsonToStruct(jsonData string, result any) error
```

## Demo

```go
// name uses two json tags
// age is defined as int, and the value of json is string
// is_use is defined as bool, the value of json is int
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
// Structure nesting and slice nesting processing
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

	jsonData2 := `{
        "name": "Bob",
        "age": "25",
        "address": {
            "city": "Shanghai",
            "country": "China"
        },
        "interests": ["reading", "swimming"]
    }`

	var person Person
	err := JsonToStruct(jsonData2, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v \n", person)
	// {Name:Bob Age:25 Address:{City:Shanghai Country:China} Interests:[reading swimming]}
}
```
