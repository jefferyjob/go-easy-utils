package cryptoUtil

import "fmt"

func ExampleMd5() {
	str := "hello"
	res := Md5(str)

	fmt.Println(res)

	// Output:
	// 5d41402abc4b2a76b9719d911017c592
}
