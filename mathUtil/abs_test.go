package mathUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbsInt64(t *testing.T) {
	testCases := []struct {
		name     string
		input    int64
		expected int64
	}{
		{
			name:     "负数转绝对值",
			input:    -12,
			expected: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Abs(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestAbsFloat32(t *testing.T) {
	testCases := []struct {
		name     string
		input    float32
		expected float32
	}{
		{
			name:     "负数转绝对值",
			input:    -12.4,
			expected: 12.4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Abs(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
