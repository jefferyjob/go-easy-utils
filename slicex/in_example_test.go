package slicex

import "fmt"

func ExampleIn() {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	res := In("c", slice)

	fmt.Println(res)

	// Output:
	// true
}
