package sliceUtil

import "fmt"

func ExampleMergeSlice() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := []int{7, 8, 9}
	res := MergeSlice(slice1, slice2, slice3)

	fmt.Printf("%+v", res)

	// Output:
	// [1 2 3 4 5 6 7 8 9]
}
