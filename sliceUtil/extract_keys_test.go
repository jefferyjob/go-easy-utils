package sliceUtil

import (
	"reflect"
	"testing"
)

func TestExtractKeys(t *testing.T) {
	type person struct {
		ID   int
		Name string
	}

	testCases := []struct {
		name        string
		inputSlices []person
		inputFunc   func(p person) int
		want        []int
	}{
		{
			name: "获取ID",
			inputSlices: []person{
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
			res := ExtractKeys(tc.inputSlices, tc.inputFunc)
			if !reflect.DeepEqual(tc.want, res) {
				t.Errorf("Expected keys %v, but got %v", tc.want, res)
			}
		})
	}
}
