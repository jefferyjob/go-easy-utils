package slicex

import "fmt"

// ExampleToMap 演示了 ToMap 函数的用法
func ExampleToMap() {
	type Person struct {
		ID   int
		Name string
	}

	people := []Person{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	peopleMap := ToMap(people, func(p Person) int {
		return p.ID
	})
	fmt.Println(peopleMap)
	// Output:
	// map[1:{1 Alice} 2:{2 Bob} 3:{3 Charlie}]
}
