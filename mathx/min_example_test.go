package mathx

import "fmt"

func ExampleMin() {
	res1 := Min([]int{5, 8, 9, 2, 1})
	res2 := Min([]float64{3.14, 2.0, -0.1, 2.2, 100.11})

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// 1
	// -0.1
}
