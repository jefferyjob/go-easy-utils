package anyx

import "fmt"

func ExampleTernary() {
	age := 15
	res := Ternary(age < 18, "未成年", "成年")

	fmt.Println(res)

	// Output:
	// 未成年
}
