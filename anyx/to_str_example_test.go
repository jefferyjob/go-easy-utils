package anyx

import "fmt"

func ExampleToStr() {
	a := 123
	res := ToStr(a)

	fmt.Println(res)

	// Output:
	// 123
}
