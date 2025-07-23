package strx

import (
	"fmt"
)

func ExampleCut() {
	fmt.Println(Cut("this is-AAA", "A"))
	fmt.Println(Cut("this is-AAA", "A", 1))

	// Output:
	// this is-
	// this is-AA
}
