package mathx

import "fmt"

func ExampleAbs() {
	res1 := Abs(-6)
	res2 := Abs(9)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// 6
	// 9
}
