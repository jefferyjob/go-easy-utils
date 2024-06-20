package sliceUtil

import "fmt"

func ExampleExtractKeys() {
	type Person struct {
		ID   int
		Name string
	}
	persons := []Person{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
	}

	keys := ExtractKeys(persons, func(p Person) int {
		return p.ID
	})
	fmt.Println(keys)
	// Output: [1 2 3]
}
