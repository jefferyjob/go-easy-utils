package emojix

import "fmt"

func ExampleEncode() {
	e := "[\\u1F60E]"
	res := Decode(e)

	fmt.Println(res)

	// Output:
	// 😎
}

func ExampleDecode() {
	e := "😂"
	res := Encode(e)

	fmt.Println(res)

	// Output:
	// [\u1f602]
}
