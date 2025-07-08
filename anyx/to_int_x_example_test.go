package anyx

import "fmt"

func ExampleToInt() {
	s1 := "100"
	s2 := "-100"
	res1, _ := ToInt(s1)
	res2, _ := ToInt(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int
	// -100,int
}

func ExampleToInt8() {
	s1 := "100"
	s2 := "-100"
	res1, _ := ToInt8(s1)
	res2, _ := ToInt8(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int8
	// -100,int8
}

func ExampleToInt16() {
	s1 := "100"
	s2 := "-100"
	res1, _ := ToInt16(s1)
	res2, _ := ToInt16(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int16
	// -100,int16
}

func ExampleToInt32() {
	s1 := "100"
	s2 := "-100"
	res1, _ := ToInt32(s1)
	res2, _ := ToInt32(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int32
	// -100,int32
}

func ExampleToInt64() {
	s1 := "100"
	s2 := "-100"
	res1, _ := ToInt64(s1)
	res2, _ := ToInt64(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int64
	// -100,int64
}
