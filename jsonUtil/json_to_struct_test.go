package jsonUtil

import (
	"encoding/json"
	"fmt"
	"testing"
)

// name uses two json tags
// age is defined as int, and the value of json is string
// is_use is defined as bool, the value of json is int
func TestDemo1(t *testing.T) {
	jsonData := `{
		"name": "make",
		"age": "22",
		"is_use": "1"
	}`

	var people1 struct {
		Name  string `json:"name,omitempty"`
		Age   int    `json:"age"`
		IsUse bool   `json:"is_use"`
	}

	var people2 struct {
		Name  string `json:"name,omitempty"`
		Age   string `json:"age"`
		IsUse bool   `json:"is_use"`
	}

	if err := JsonToStruct(jsonData, &people1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("people1:%#v \n", people1)
	// return
	// people1:struct { Name string "json:\"name,omitempty\""; Age int "json:\"age\""; IsUse bool "json:\"is_use\"" }{Name:"make", Age:22, IsUse:true}

	if err := JsonToStruct(jsonData, &people2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("people2:%#v \n", people2)
	// return
	// people2:struct { Name string "json:\"name,omitempty\""; Age string "json:\"age\""; IsUse bool "json:\"is_use\"" }{Name:"make", Age:"22", IsUse:true}
}

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

func BenchmarkJsonUnmarshal(b *testing.B) {
	jsonData := `{
		"name": "make",
		"age": 22,
		"is_use": true
	}`
	var people struct {
		Name  string `json:"name,omitempty"`
		Age   int    `json:"age"`
		IsUse bool   `json:"is_use"`
	}
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonData), &people)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}
}

func BenchmarkJsonToStruct(b *testing.B) {
	jsonData := `{
		"name": "make",
		"age": "22",
		"is_use": true
	}`
	var people struct {
		Name  string `json:"name,omitempty"`
		Age   int    `json:"age"`
		IsUse bool   `json:"is_use"`
	}
	for i := 0; i < b.N; i++ {
		err := JsonToStruct(jsonData, &people)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}
}

// 非法验证：非法的json字符串
func TestJsonToStructErrJson(t *testing.T) {
	jsonData := `{"name":}`
	type People struct {
		Name string `json:"name"`
	}
	var people People
	err := JsonToStruct(jsonData, &people)
	if err == nil {
		t.Errorf("err %s", err)
		return
	}
}

// 非法验证：非指针的result
func TestJsonToStructErrResult(t *testing.T) {
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
	err := JsonToStruct(jsonData, people)
	if err == nil {
		t.Errorf("err %s", err)
		return
	}
}

// 合法验证：空数据
func TestJsonToStructEmptyValue(t *testing.T) {
	jsonData := `{
		"name": "make",
		"age": "",
		"source": "",
		"num": "",
		"status": ""
	}`
	var people struct {
		Name   string  `json:"name,omitempty"`
		Age    int     `json:"age"`
		Source uint    `json:"source"`
		Num    float64 `json:"num"`
		Status bool    `json:"status"`
	}
	err := JsonToStruct(jsonData, &people)
	if err != nil {
		t.Errorf("err %s", err)
		return
	}
}

// 非法验证：非法的int
func TestJsonToStructErrInt(t *testing.T) {
	jsonData := `{
		"name": "make",
		"age": "test abc"
	}`
	var people struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age"`
	}
	err := JsonToStruct(jsonData, people)
	if err == nil {
		t.Errorf("err %s", err)
		return
	}
}

// 非法验证：非法的uint
func TestJsonToStructErrUint(t *testing.T) {
	jsonData := `{
		"name": "make",
		"age": "test abc"
	}`
	var people struct {
		Name string `json:"name,omitempty"`
		Age  uint   `json:"age"`
	}
	err := JsonToStruct(jsonData, people)
	if err == nil {
		t.Errorf("err %s", err)
		return
	}
}

// 非法验证：非法的float
func TestJsonToStructErrFloat(t *testing.T) {
	jsonData := `{
		"name": "make",
		"age": "test abc"
	}`
	var people struct {
		Name string  `json:"name,omitempty"`
		Age  float64 `json:"age"`
	}
	err := JsonToStruct(jsonData, people)
	if err == nil {
		t.Errorf("err %s", err)
		return
	}
}

// 非法验证：嵌套的数据类型错误
func TestJsonToStructNestErrInt(t *testing.T) {
	jsonData := `{
        "name": "John Doe",
        "address": {
            "number": "test abc"
        }
    }`
	type Address struct {
		Number int `json:"number"`
	}
	type Person struct {
		Name    string  `json:"name"`
		Address Address `json:"address"`
	}
	people := &Person{}
	err := JsonToStruct(jsonData, people)
	if err == nil {
		t.Errorf("err %s", err)
		return
	}
}

// 合法验证：多层级嵌套
func TestJsonToStructMoreNest(t *testing.T) {
	type Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	}

	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Source  float64  `json:"source"`
		Number  int      `json:"number"`
		Status  bool     `json:"status"`
		Address Address  `json:"address"`
		Emails  []string `json:"emails"`
	}

	jsonData := `{
        "name": "John Doe",
        "age": 30,
 		"source": "99.99",
		"status": true,
        "address": {
            "city": "Shanghai",
            "country": "China"
        },
        "emails": [
            "john.doe@example.com",
            "jdoe@example.com"
        ]
    }`

	expectedPerson := Person{
		Name:   "John Doe",
		Age:    30,
		Status: true,
		Address: Address{
			City:    "Shanghai",
			Country: "China",
		},
		Emails: []string{"john.doe@example.com", "jdoe@example.com"},
	}

	var resultPerson Person
	err := JsonToStruct(jsonData, &resultPerson)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resultPerson.Name != expectedPerson.Name {
		t.Errorf("Name field mismatch: expected %s but got %s", expectedPerson.Name, resultPerson.Name)
	}

	if resultPerson.Age != expectedPerson.Age {
		t.Errorf("Age field mismatch: expected %d but got %d", expectedPerson.Age, resultPerson.Age)
	}

	if resultPerson.Status != expectedPerson.Status {
		t.Errorf("Status field mismatch: expected %v but got %v", expectedPerson.Status, resultPerson.Status)
	}

	if resultPerson.Address.City != expectedPerson.Address.City {
		t.Errorf("Address.City field mismatch: expected %s but got %s", expectedPerson.Address.City, resultPerson.Address.City)
	}

	if resultPerson.Address.Country != expectedPerson.Address.Country {
		t.Errorf("Address.Country field mismatch: expected %s but got %s", expectedPerson.Address.Country, resultPerson.Address.Country)
	}

	if len(resultPerson.Emails) != len(expectedPerson.Emails) {
		t.Errorf("Emails length mismatch: expected %d but got %d", len(expectedPerson.Emails), len(resultPerson.Emails))
	}

	for i, expectedEmail := range expectedPerson.Emails {
		if resultPerson.Emails[i] != expectedEmail {
			t.Errorf("Emails[%d] mismatch: expected %s but got %s", i, expectedEmail, resultPerson.Emails[i])
		}
	}
}

// 多层级json测试
//func TestJsonToStruct4(t *testing.T) {
//	type Address struct {
//		City    string `json:"city"`
//		Street  string `json:"street"`
//		Zipcode uint64 `json:"zipcode"`
//	}
//
//	type Score struct {
//		Subject string `json:"subject"`
//		Score   int    `json:"score"`
//	}
//
//	type Student struct {
//		Name    string  `json:"name,omitempty"`
//		Age     int     `json:"age,omitempty"`
//		Address Address `json:"address"`
//		Scores  []Score `json:"scores,omitempty"`
//	}
//	jsonStr4 := `{
//		"name": "Alice",
//		"age": 30,
//		"address": {
//			"city": "Beijing",
//			"street": "Zhangsan Street",
//			"zipcode": 100
//		},
//		"scores": [
//			{"subject": "Math", "score": 80},
//			{"subject": "English", "score": 90}
//		]
//	}`
//
//	var student Student
//	if err := JsonToStruct(jsonStr4, &student); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Printf("%+v \n", student)
//}

//func TestJsonToStruct5(t *testing.T) {
//	type Student struct {
//		Name string `json:"name,omitempty"`
//		Age  int    `json:"age,omitempty"`
//	}
//	jsonStr4 := `{
//		"name":null,
//		"age": "30"
//	}`
//
//	var student Student
//	if err := JsonToStruct(jsonStr4, &student); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Printf("%+v \n", student)
//}

//func TestJsonToStruct6(t *testing.T) {
//	type Student struct {
//		Name any `json:"name,omitempty"`
//		Age  int         `json:"age,omitempty"`
//	}
//	jsonStr4 := `{
//		"name":"zhangsan",
//		"age": "123"
//	}`
//
//	var student Student
//	if err := JsonToStruct(jsonStr4, &student); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Printf("%+v \n", student)
//}

//func TestJsonToStruct7(t *testing.T) {
//	type Student struct {
//		Name  bool `json:"name"`
//		Name2 uint `json:"name2"`
//		Name3 uint `json:"name3"`
//		Age   int  `json:"age"`
//	}
//	jsonStr4 := `{
//		"name": true,
//		"name2": -1,
//		"name3": null,
//		"age": "123"
//	}`
//
//	var student Student
//	if err := JsonToStruct(jsonStr4, &student); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Printf("%#v \n", student)
//}
