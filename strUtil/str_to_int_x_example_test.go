package strUtil

import "fmt"

func ExampleStrToInt() {
	s1 := "100"
	s2 := "-100"
	res1 := StrToInt(s1)
	res2 := StrToInt(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int
	// -100,int
}

func ExampleStrToInt8() {
	s1 := "100"
	s2 := "-100"
	res1 := StrToInt8(s1)
	res2 := StrToInt8(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int8
	// -100,int8
}

func ExampleStrToInt16() {
	s1 := "100"
	s2 := "-100"
	res1 := StrToInt16(s1)
	res2 := StrToInt16(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int16
	// -100,int16
}

func ExampleStrToInt32() {
	s1 := "100"
	s2 := "-100"
	res1 := StrToInt32(s1)
	res2 := StrToInt32(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int32
	// -100,int32
}

func ExampleStrToInt64() {
	s1 := "100"
	s2 := "-100"
	res1 := StrToInt64(s1)
	res2 := StrToInt64(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int64
	// -100,int64
}
