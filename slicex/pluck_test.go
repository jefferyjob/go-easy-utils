package slicex

import (
	"reflect"
	"testing"
)

func TestPluck(t *testing.T) {
	type person struct {
		ID   int
		Name string
	}

	testCases := []struct {
		name      string
		inputs    []person
		inputFunc func(p person) int
		want      []int
	}{
		{
			name: "获取ID",
			inputs: []person{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			inputFunc: func(p person) int {
				return p.ID
			},
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Pluck(tc.inputs, tc.inputFunc)
			if !reflect.DeepEqual(tc.want, res) {
				t.Errorf("Expected keys %v, but got %v", tc.want, res)
			}
		})
	}
}
