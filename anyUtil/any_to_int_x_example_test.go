package anyUtil

import "fmt"

func ExampleAnyToInt() {
	s1 := "100"
	s2 := "-100"
	res1, _ := AnyToInt(s1)
	res2, _ := AnyToInt(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int
	// -100,int
}

func ExampleAnyToInt8() {
	s1 := "100"
	s2 := "-100"
	res1, _ := AnyToInt8(s1)
	res2, _ := AnyToInt8(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int8
	// -100,int8
}

func ExampleAnyToInt16() {
	s1 := "100"
	s2 := "-100"
	res1, _ := AnyToInt16(s1)
	res2, _ := AnyToInt16(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int16
	// -100,int16
}

func ExampleAnyToInt32() {
	s1 := "100"
	s2 := "-100"
	res1, _ := AnyToInt32(s1)
	res2, _ := AnyToInt32(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int32
	// -100,int32
}

func ExampleAnyToInt64() {
	s1 := "100"
	s2 := "-100"
	res1, _ := AnyToInt64(s1)
	res2, _ := AnyToInt64(s2)

	fmt.Printf("%d,%T\n", res1, res1)
	fmt.Printf("%d,%T\n", res2, res2)

	// Output:
	// 100,int64
	// -100,int64
}
