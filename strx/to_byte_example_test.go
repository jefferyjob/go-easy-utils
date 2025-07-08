package strx

import "fmt"

func ExampleToBytes() {
	s := "hello"
	res := ToBytes(s)

	fmt.Printf("%v", res)

	// Output:
	// [104 101 108 108 111]
}
