package slicex

import "fmt"

func ExampleChunk() {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	res := Chunk(slice, 3)

	fmt.Printf("%+v", res)

	// Output:
	// [[a b c] [d e f] [g h i] [j k]]
}
