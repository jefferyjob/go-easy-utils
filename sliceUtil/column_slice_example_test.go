package sliceUtil

import "fmt"

func ExampleColumnSlice() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 18},
		{"Bob", 20},
		{"Charlie", 22},
	}

	res := ColumnSlice(people, "Age")

	fmt.Printf("%+v", res)

	// Output:
	// [18 20 22]
}
