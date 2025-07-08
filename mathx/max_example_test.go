package mathx

import "fmt"

func ExampleMax() {
	res1 := Max([]int{5, 8, 9, 2, 1})
	res2 := Max([]float64{3.14, 2.0, -0.1, 2.2, 100.11})

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// 9
	// 100.11
}
