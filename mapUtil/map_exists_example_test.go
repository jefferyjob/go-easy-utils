package mapUtil

import "fmt"

func ExampleMapKeyExists() {
	m := map[string]string{
		"foo": "bar",
		"baz": "123",
	}
	res := MapKeyExists(m, "foo")

	fmt.Println(res)

	// Output:
	// true
}

func ExampleMapValueExists() {
	m := map[string]string{
		"foo": "bar",
		"baz": "123",
	}
	res := MapValueExists(m, "123")

	fmt.Println(res)

	// Output:
	// true
}
