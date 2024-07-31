package sliceUtil

import (
	"reflect"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	type person struct {
		ID   int
		Name string
	}

	testCases := []struct {
		name        string
		inputSlices []person
		inputFunc   func(p person) int
		want        map[int]person
	}{
		{
			name: "以id为map的key",
			inputSlices: []person{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 3, Name: "Charlie"},
			},
			inputFunc: func(p person) int {
				return p.ID
			},
			want: map[int]person{
				1: {ID: 1, Name: "Alice"},
				2: {ID: 2, Name: "Bob"},
				3: {ID: 3, Name: "Charlie"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := SliceToMap(tc.inputSlices, tc.inputFunc)
			if !reflect.DeepEqual(tc.want, res) {
				t.Errorf("failed. Expected %v, got %v", tc.want, res)
			}
		})
	}
}
