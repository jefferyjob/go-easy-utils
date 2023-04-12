package strUtil

import "fmt"

func ExampleStrToUint() {
	s := "200"
	res := StrToUint(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint
}

func ExampleStrToUint8() {
	s := "200"
	res := StrToUint8(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint8
}

func ExampleStrToUint16() {
	s := "200"
	res := StrToUint16(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint16
}

func ExampleStrToUint32() {
	s := "200"
	res := StrToUint32(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint32
}

func ExampleStrToUint64() {
	s := "200"
	res := StrToUint64(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint64
}
