package slicex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIs(t *testing.T) {
	testCases := []struct {
		name  string
		input any
		want  bool
	}{
		{
			name:  "slice of int",
			input: []int{1, 1, 3},
			want:  true,
		},
		{
			name:  "slice of any",
			input: []any{1, 2, "a"},
			want:  true,
		},
		{
			name: "slice of map",
			input: []map[string]any{
				{"1": 1},
				{"c": 89},
			},
			want: true,
		},
		{
			name:  "string",
			input: "1234",
			want:  false,
		},
		{
			name:  "channel",
			input: make(chan int),
			want:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Is(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
