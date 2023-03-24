package jsonUtil

import (
	"fmt"
	"log"
	"testing"
)

func TestJsonToStruct1(t *testing.T) {
	type Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	}

	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Address Address  `json:"address"`
		Emails  []string `json:"emails"`
	}

	jsonData := `{
        "name": "John Doe",
        "age": 30,
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
		Name: "John Doe",
		Age:  30,
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

func TestJsonToStruct2(t *testing.T) {
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

	jsonData1 := `{
        "name": "Bob",
        "age": 25,
        "address": {
            "city": "Shanghai",
            "country": "China"
        },
        "interests": ["reading", "swimming"]
    }`

	var person1 Person
	err := JsonToStruct(jsonData1, &person1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("person1：%+v \n", person1)

}

func TestJsonToStruct3(t *testing.T) {
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

	var person2 Person
	err := JsonToStruct(jsonData2, &person2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("person2：%+v，， address：%+v \n", person2, person2.Address)
}

// 多层级json测试
func TestJsonToStruct4(t *testing.T) {
	type Address struct {
		City    string `json:"city"`
		Street  string `json:"street"`
		Zipcode uint64 `json:"zipcode"`
	}

	type Score struct {
		Subject string `json:"subject"`
		Score   int    `json:"score"`
	}

	type Student struct {
		Name    string  `json:"name,omitempty"`
		Age     int     `json:"age,omitempty"`
		Address Address `json:"address"`
		Scores  []Score `json:"scores,omitempty"`
	}
	jsonStr4 := `{
		"name": "Alice",
		"age": 30,
		"address": {
			"city": "Beijing",
			"street": "Zhangsan Street",
			"zipcode": 100
		},
		"scores": [
			{"subject": "Math", "score": 80},
			{"subject": "English", "score": 90}
		]
	}`

	var student Student
	if err := JsonToStruct(jsonStr4, &student); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v \n", student)
}

func TestJsonToStruct5(t *testing.T) {
	type Student struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}
	jsonStr4 := `{
		"name":null,
		"age": "30"
	}`

	var student Student
	if err := JsonToStruct(jsonStr4, &student); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v \n", student)
}

func TestJsonToStruct6(t *testing.T) {
	type Student struct {
		Name interface{} `json:"name,omitempty"`
		Age  int         `json:"age,omitempty"`
	}
	jsonStr4 := `{
		"name":"zhangsan",
		"age": "123"
	}`

	var student Student
	if err := JsonToStruct(jsonStr4, &student); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v \n", student)
}

func TestJsonToStruct7(t *testing.T) {
	type Student struct {
		Name  bool `json:"name"`
		Name2 uint `json:"name2"`
		Name3 uint `json:"name3"`
		Age   int  `json:"age"`
	}
	jsonStr4 := `{
		"name": true,
		"name2": -1,
		"name3": null,
		"age": "123"
	}`

	var student Student
	if err := JsonToStruct(jsonStr4, &student); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v \n", student)
}

func TestToUint64(t *testing.T) {
	fmt.Println(toUint64(""))
}
