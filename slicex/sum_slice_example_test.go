package slicex

import "fmt"

func ExampleSum() {
	slice := []int{1, 2, 3, 4, 5}
	res := Sum(slice)

	fmt.Println(res)

	// Output:
	// 15
}
