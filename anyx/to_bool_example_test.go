package anyx

import "fmt"

func ExampleToBool() {
	a1 := ""
	a2 := "hello"
	res1 := ToBool(a1)
	res2 := ToBool(a2)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// false
	// true
}
