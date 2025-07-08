package anyx

import "fmt"

func ExampleToUint() {
	s := "200"
	res, _ := ToUint(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint
}

func ExampleToUint8() {
	s := "200"
	res, _ := ToUint8(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint8
}

func ExampleToUint16() {
	s := "200"
	res, _ := ToUint16(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint16
}

func ExampleToUint32() {
	s := "200"
	res, _ := ToUint32(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint32
}

func ExampleToUint64() {
	s := "200"
	res, _ := ToUint64(s)

	fmt.Printf("%d,%T\n", res, res)

	// Output:
	// 200,uint64
}
