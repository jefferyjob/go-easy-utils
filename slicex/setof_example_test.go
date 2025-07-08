package slicex

import "fmt"

func ExampleDiff() {
	s1 := []int64{1, 2, 5, 7}
	s2 := []int64{5, 6, 7}

	diff := Diff(s1, s2)
	fmt.Println(diff)

	// Output:
	// [1 2]
}

// 演示 SymmetricDiff 函数的用法
func ExampleSymmetricDiff() {
	s1 := []int64{1, 5, 7}
	s2 := []int64{5, 6, 7, 8, 9}
	s3 := []int64{9, 10, 11}

	diff := SymmetricDiff(s1, s2, s3)
	fmt.Println(diff)

	// Output:
	// [1 6 8 10 11]
}

// 演示 ExampleIntersect 函数的用法
func ExampleIntersect() {
	s1 := []int64{1, 5, 7}
	s2 := []int64{5, 6, 7, 8, 9}

	diff := Intersect(s1, s2)
	fmt.Println(diff)

	// Output:
	// [5 7]
}
