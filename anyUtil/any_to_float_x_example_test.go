package anyUtil

import "fmt"

func ExampleAnyToFloat32() {
	f := "3"
	res, _ := AnyToFloat32(f)

	fmt.Printf("%T\n", res)

	// Output:
	// float32
}

func ExampleAnyToFloat64() {
	f := "5"
	res, _ := AnyToFloat64(f)

	fmt.Printf("%T\n", res)

	// Output:
	// float64
}
