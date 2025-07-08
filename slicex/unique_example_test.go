package slicex

import "fmt"

func ExampleUnique() {
	slice := []int{1, 2, 3, 2, 4, 4, 5}
	res := Unique(slice)

	fmt.Printf("%+v", res)

	// Output:
	// [1 2 3 4 5]
}
