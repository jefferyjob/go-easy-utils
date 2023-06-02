# jsonUtil

## Introduce

JsonToStruct converts JSON data to a Go structure.
Parameters:
- jsonData: A string containing the JSON data.
- val: A pointer to the structure variable to be filled.

Returns:
- error: If conversion fails or an error occurs, the corresponding error is returned. If successful, nil is returned.

Functionality:
- Checks if the val parameter is a non-nil pointer type, returning ErrPoint if it is not. 
- Parses jsonData into a map[string]any variable called data. 
- Retrieves the value and type of the structure pointed to by val using reflection. 
- Iterates through the fields of the structure:
- Retrieves the field's type, name, and value. 
- Gets the JSON tag for the field. 
- Performs the appropriate handling if a key-value pair corresponding to the JSON tag exists in data:
- If the field is a primitive type (string, integer, float, boolean), parses it into the corresponding type's value.
- If the field is a struct type, recursively calls the JsonToStruct function to convert the sub-structure to JSON.
- If the field is a map type, uses the parseMap function to convert the sub-map to JSON.
- If the field is a slice type, uses the parseSlice function to convert the sub-slice to JSON.
- If the field is an interface type, sets the value to nil or the corresponding value.

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/v2/jsonUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/jsonUtil"
)
```

## Functions

```go
func JsonToStruct(jsonData string, val any) error
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
