package sliceUtil

import "fmt"

func ExampleSumSlice() {
	slice := []int{1, 2, 3, 4, 5}
	res := SumSlice(slice)

	fmt.Println(res)

	// Output:
	// 15
}
