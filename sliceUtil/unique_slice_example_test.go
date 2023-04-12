package sliceUtil

import "fmt"

func ExampleUniqueSlice() {
	slice := []int{1, 2, 3, 2, 4, 4, 5}
	res := UniqueSlice(slice)

	fmt.Printf("%+v", res)

	// Output:
	// [1 2 3 4 5]
}
