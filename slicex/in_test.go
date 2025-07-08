package slicex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInStrs(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []string
		element  string
		expected bool
	}{
		{
			name:     "apple in slice",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "apple",
			expected: true,
		},
		{
			name:     "banana in slice",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "banana",
			expected: true,
		},
		{
			name:     "orange not in slice",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "orange",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestInInts(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int
		element  int
		expected bool
	}{
		{
			name:     "3 in slice",
			slice:    []int{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "6 not in slice",
			slice:    []int{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestInInt8s(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int8
		element  int8
		expected bool
	}{
		{
			name:     "3 in slice",
			slice:    []int8{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "6 not in slice",
			slice:    []int8{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestInInt16s(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int16
		element  int16
		expected bool
	}{
		{
			name:     "3 in slice",
			slice:    []int16{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "6 not in slice",
			slice:    []int16{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestInInt32s(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int32
		element  int32
		expected bool
	}{
		{
			name:     "3 in slice",
			slice:    []int32{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "6 not in slice",
			slice:    []int32{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestInInt64s(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int64
		element  int64
		expected bool
	}{
		{
			name:     "3 in slice",
			slice:    []int64{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "6 not in slice",
			slice:    []int64{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestInUints(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []uint
		element  uint
		expected bool
	}{
		{
			name:     "3 in slice",
			slice:    []uint{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "6 not in slice",
			slice:    []uint{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestInUint8s(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []uint8
		element  uint8
		expected bool
	}{
		{
			name:     "3 in slice",
			slice:    []uint8{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "6 not in slice",
			slice:    []uint8{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestInUint16s(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []uint16
		element  uint16
		expected bool
	}{
		{
			name:     "3 in slice",
			slice:    []uint16{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "6 not in slice",
			slice:    []uint16{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestInUint32s(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []uint32
		element  uint32
		expected bool
	}{
		{
			name:     "3 in slice",
			slice:    []uint32{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "6 not in slice",
			slice:    []uint32{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestInUint64s(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []uint64
		element  uint64
		expected bool
	}{
		{
			name:     "3 in slice",
			slice:    []uint64{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "6 not in slice",
			slice:    []uint64{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := In(tc.element, tc.slice)
			assert.Equal(t, tc.expected, res)
		})
	}
}
