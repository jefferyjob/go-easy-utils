package anyUtil

import "fmt"

func ExampleAnyToUint() {
	s := "200"
	res, _ := AnyToUint(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint
}

func ExampleAnyToUint8() {
	s := "200"
	res, _ := AnyToUint8(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint8
}

func ExampleAnyToUint16() {
	s := "200"
	res, _ := AnyToUint16(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint16
}

func ExampleAnyToUint32() {
	s := "200"
	res, _ := AnyToUint32(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint32
}

func ExampleAnyToUint64() {
	s := "200"
	res, _ := AnyToUint64(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint64
}
