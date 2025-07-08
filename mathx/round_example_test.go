package mathx

import "fmt"

func ExampleRound() {
	res1 := Round(3.14)
	res2 := Round(3.56)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// 3
	// 4
}
