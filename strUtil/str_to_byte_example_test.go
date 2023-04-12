package strUtil

import "fmt"

func ExampleStrToBytes() {
	s := "hello"
	res := StrToBytes(s)

	fmt.Printf("%v", res)

	// Output:
	// [104 101 108 108 111]
}
