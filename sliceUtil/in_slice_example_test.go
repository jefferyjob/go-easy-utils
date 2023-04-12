package sliceUtil

import "fmt"

func ExampleInSlice() {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	res := InSlice("c", slice)

	fmt.Println(res)

	// Output:
	// true
}
