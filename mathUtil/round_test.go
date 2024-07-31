package mathUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoundFloat32(t *testing.T) {
	testCases := []struct {
		name     string
		input    float32
		expected int
	}{
		{
			name:     "负数四舍五入",
			input:    -12.4,
			expected: -12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Round(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestRoundFloat64(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected int
	}{
		{
			name:     "正数四舍五入",
			input:    12.5,
			expected: 13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Round(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
