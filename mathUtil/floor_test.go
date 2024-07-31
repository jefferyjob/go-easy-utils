package mathUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloorFloat32(t *testing.T) {
	testCases := []struct {
		name     string
		input    float32
		expected int
	}{
		{
			name:     "负浮点数取整",
			input:    -12.4,
			expected: -13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Floor(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestFloorFloat64(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected int
	}{
		{
			name:     "正浮点数取整",
			input:    12.4,
			expected: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Floor(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
