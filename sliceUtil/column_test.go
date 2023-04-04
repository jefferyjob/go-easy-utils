package sliceUtil

import (
	"reflect"
	"testing"
)

//func TestColumnSliceDemo(t *testing.T) {
//	type Person struct {
//		Name string
//		Age  int
//	}
//	people := []Person{
//		{"Alice", 18},
//		{"Bob", 20},
//		{"Charlie", 22},
//	}
//
//	// 获取年龄列
//	ages := Column(people, "Age")
//	fmt.Println(ages) // 输出：[18 20 22]
//}

func TestColumnSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name   string
		input  []Person
		column string
		output []interface{}
	}{
		{
			name:   "success",
			input:  []Person{{"Alice", 18}, {"Bob", 20}, {"Charlie", 22}},
			column: "Age",
			output: []any{18, 20, 22},
		},
		{
			name:   "column not found",
			input:  []Person{{"Alice", 18}, {"Bob", 20}, {"Charlie", 22}},
			column: "Gender",
			output: []any{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Column(tc.input, tc.column)
			if !reflect.DeepEqual(output, tc.output) {
				t.Errorf("expected %v, but got %v", tc.output, output)
			}
		})
	}
}
