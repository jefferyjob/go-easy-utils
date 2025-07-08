package anyx

import "fmt"

func ExampleToFloat32() {
	f := "3"
	res, _ := ToFloat32(f)

	fmt.Printf("%T\n", res)

	// Output:
	// float32
}

func ExampleToFloat64() {
	f := "5"
	res, _ := ToFloat64(f)

	fmt.Printf("%T\n", res)

	// Output:
	// float64
}
