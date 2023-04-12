package sliceUtil

import "fmt"

func ExampleChunkSlice() {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	res := ChunkSlice(slice, 3)

	fmt.Printf("%+v", res)

	// Output:
	// [[a b c] [d e f] [g h i] [j k]]
}
