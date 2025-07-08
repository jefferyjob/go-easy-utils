package mathx

import "fmt"

func ExampleCeil() {
	res1 := Ceil(3.14)
	res2 := Ceil(-6.5)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// 4
	// -6
}
