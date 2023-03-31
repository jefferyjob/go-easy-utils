package sliceUtil

import (
	"errors"
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
//	ages, err := ColumnSlice(people, "Age")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(ages) // 输出：[18 20 22]
//}

func TestColumnSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name   string
		input  interface{}
		column string
		output []interface{}
		err    error
	}{
		{
			name:   "success",
			input:  []Person{{"Alice", 18}, {"Bob", 20}, {"Charlie", 22}},
			column: "Age",
			output: []interface{}{18, 20, 22},
			err:    nil,
		},
		{
			name:   "not a slice",
			input:  Person{"Alice", 18},
			column: "Age",
			output: nil,
			err:    errors.New("not a slice"),
		},
		{
			name:   "column not found",
			input:  []Person{{"Alice", 18}, {"Bob", 20}, {"Charlie", 22}},
			column: "Gender",
			output: nil,
			err:    errors.New("column not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := ColumnSlice(tc.input, tc.column)
			if !reflect.DeepEqual(output, tc.output) {
				t.Errorf("expected %v, but got %v", tc.output, output)
			}
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("expected %v, but got %v", tc.err, err)
			}
		})
	}
}
