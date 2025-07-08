package slicex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColumn(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name        string
		input       []Person
		inputColumn string
		want        []any
	}{
		{
			name:        "success",
			input:       []Person{{"Alice", 18}, {"Bob", 20}, {"Charlie", 22}},
			inputColumn: "Age",
			want:        []any{18, 20, 22},
		},
		{
			name:        "column not found",
			input:       []Person{{"Alice", 18}, {"Bob", 20}, {"Charlie", 22}},
			inputColumn: "Gender",
			want:        []any{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Column(tc.input, tc.inputColumn)
			assert.Equal(t, tc.want, res)
		})
	}
}
