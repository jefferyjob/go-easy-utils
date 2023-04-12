package byteUtil

import "fmt"

func ExampleBytesToStr() {
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	res := BytesToStr(b)

	fmt.Println(res)

	// Output:
	// hello
}
