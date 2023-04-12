package anyUtil

import "fmt"

func ExampleAnyToStr() {
	a := 123
	res := AnyToStr(a)

	fmt.Println(res)

	// Output:
	// 123
}
