package mapx

import "fmt"

func ExampleKeyExists() {
	m := map[string]string{
		"foo": "bar",
		"baz": "123",
	}
	res := KeyExists(m, "foo")

	fmt.Println(res)

	// Output:
	// true
}

func ExampleValueExists() {
	m := map[string]string{
		"foo": "bar",
		"baz": "123",
	}
	res := ValueExists(m, "123")

	fmt.Println(res)

	// Output:
	// true
}
