package mathUtil

import "fmt"

func ExampleFloor() {
	res1 := Floor(3.14)
	res2 := Floor(-6.5)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// 3
	// -7
}
