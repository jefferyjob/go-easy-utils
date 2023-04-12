package jsonUtil

import "fmt"

func ExampleJsonToStruct() {
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
	fmt.Printf("%+v", people)

	// Output:
	// {Name:make Age:22 IsUse:true}
}
