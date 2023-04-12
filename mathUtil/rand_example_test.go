package mathUtil

import "fmt"

func ExampleRand() {
	minInt := 1
	maxInt := 100
	res1 := Rand(minInt, maxInt)
	if res1 >= 1 && res1 < 100 {
		fmt.Println("ok")
	}

	minFloat := 1.0
	maxFloat := 10.0
	res2 := Rand(minFloat, maxFloat)
	if res2 >= 1.0 && res2 < 10.0 {
		fmt.Println("ok")
	}

	// Output:
	// ok
	// ok
}
