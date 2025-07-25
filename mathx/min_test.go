package mathx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "空切片",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "普通切片",
			input:    []int{3, 7, 1, 9, 3, 0, 2, 2},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Min(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
