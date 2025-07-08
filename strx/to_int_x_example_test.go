package strx

import "fmt"

func ExampleToInt() {
	s1 := "100"
	s2 := "-100"
	res1 := ToInt(s1)
	res2 := ToInt(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int
	// -100,int
}

func ExampleToInt8() {
	s1 := "100"
	s2 := "-100"
	res1 := ToInt8(s1)
	res2 := ToInt8(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int8
	// -100,int8
}

func ExampleToInt16() {
	s1 := "100"
	s2 := "-100"
	res1 := ToInt16(s1)
	res2 := ToInt16(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int16
	// -100,int16
}

func ExampleToInt32() {
	s1 := "100"
	s2 := "-100"
	res1 := ToInt32(s1)
	res2 := ToInt32(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int32
	// -100,int32
}

func ExampleToInt64() {
	s1 := "100"
	s2 := "-100"
	res1 := ToInt64(s1)
	res2 := ToInt64(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int64
	// -100,int64
}
