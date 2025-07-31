package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		filterFn func(int) bool
		want     []int
	}{
		{
			name:     "filter even numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			filterFn: func(i int) bool { return i%2 == 0 },
			want:     []int{2, 4, 6},
		},
		{
			name:     "filter all",
			input:    []int{1, 2, 3},
			filterFn: func(i int) bool { return true },
			want:     []int{1, 2, 3},
		},
		{
			name:     "filter none",
			input:    []int{1, 2, 3},
			filterFn: func(i int) bool { return false },
			want:     []int{},
		},
		{
			name:     "empty input",
			input:    []int{},
			filterFn: func(i int) bool { return i > 0 },
			want:     []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Filter(tc.input, tc.filterFn)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFilterStrings(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		filterFn func(string) bool
		want     []string
	}{
		{
			name:     "length > 2",
			input:    []string{"go", "java", "c", "rust"},
			filterFn: func(s string) bool { return len(s) > 2 },
			want:     []string{"java", "rust"},
		},
		{
			name:     "empty input",
			input:    []string{},
			filterFn: func(s string) bool { return len(s) > 0 },
			want:     []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Filter(tc.input, tc.filterFn)
			assert.Equal(t, tc.want, got)
		})
	}
}
