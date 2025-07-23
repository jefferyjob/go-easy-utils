package anyx

import "fmt"

func ExampleTernary() {
	age := 15
	res1 := Ternary(age < 18, "未成年", "成年")
	res2 := Ternary(true, "Yes", "No")
	res3 := Ternary(false, "Yes", "No")

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	// Output:
	// 未成年
	// Yes
	// No
}
