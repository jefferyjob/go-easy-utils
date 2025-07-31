package slicex

import (
	"fmt"
)

func ExampleFilter() {
	numbers := []int{1, 2, 3, 4, 5}
	even := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(even)

	words := []string{"go", "java", "c", "python"}
	longWords := Filter(words, func(s string) bool {
		return len(s) > 2
	})
	fmt.Println(longWords)

	// Output:
	// [2 4]
	// [java python]
}
