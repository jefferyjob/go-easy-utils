package floatx

import "fmt"

func ExampleFloat32ToStr() {
	var f float32 = 3.1
	res := Float32ToStr(f)

	fmt.Println(res)

	// Output:
	// 3.1
}

func ExampleFloat32ToFloat64() {
	var f float32 = 3.1
	res := Float32ToFloat64(f)

	fmt.Println(res)

	// Output:
	// 3.1
}

func ExampleFloat64ToStr() {
	f := 3.14
	res := Float64ToStr(f)

	fmt.Println(res)

	// Output:
	// 3.14
}

func ExampleFloat64ToFloat32() {
	f := 3.14
	res := Float64ToFloat32(f)

	fmt.Println(res)

	// Output:
	// 3.14
}
