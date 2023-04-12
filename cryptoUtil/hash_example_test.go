package cryptoUtil

import "fmt"

func ExampleHashSHA256() {
	str := "hello"
	res := HashSHA256(str)

	fmt.Println(res)

	// Output:
	// 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}
