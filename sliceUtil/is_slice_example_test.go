package sliceUtil

import "fmt"

func ExampleIsSlice() {
	slice1 := []int{1, 2, 3}
	slice2 := make(chan int)

	res1 := IsSlice(slice1)
	res2 := IsSlice(slice2)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// true
	// false
}
