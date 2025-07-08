package bytex

import "fmt"

func ExampleToStr() {
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	res := ToStr(b)

	fmt.Println(res)

	// Output:
	// hello
}
