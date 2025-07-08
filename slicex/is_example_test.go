package slicex

import "fmt"

func ExampleIs() {
	slice1 := []int{1, 2, 3}
	slice2 := make(chan int)

	res1 := Is(slice1)
	res2 := Is(slice2)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// true
	// false
}
