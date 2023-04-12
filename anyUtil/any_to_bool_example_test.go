package anyUtil

import "fmt"

func ExampleAnyToBool() {
	a1 := ""
	a2 := "hello"
	res1 := AnyToBool(a1)
	res2 := AnyToBool(a2)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// false
	// true
}
