package slicex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerges(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]any
		expected []any
	}{
		{
			name: "merge multiple slices",
			slices: [][]any{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []any{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMergeStrs(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]string
		expected []string
	}{
		{
			name: "merge multiple string slices",
			slices: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
				{"g", "h", "i"},
			},
			expected: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMergeInts(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			name: "merge multiple int slices",
			slices: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMergeInt8s(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]int8
		expected []int8
	}{
		{
			name: "merge multiple int8 slices",
			slices: [][]int8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMergeInt16s(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]int16
		expected []int16
	}{
		{
			name: "merge multiple int16 slices",
			slices: [][]int16{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int16{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMergeInt32s(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]int32
		expected []int32
	}{
		{
			name: "merge multiple int32 slices",
			slices: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}
func TestMergeInt64s(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]int64
		expected []int64
	}{
		{
			name: "merge multiple int64 slices",
			slices: [][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMergeUints(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]uint
		expected []uint
	}{
		{
			name: "merge multiple uint slices",
			slices: [][]uint{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []uint{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMergeUint8s(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]uint8
		expected []uint8
	}{
		{
			name: "merge multiple uint8 slices",
			slices: [][]uint8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMergeUint16s(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]uint16
		expected []uint16
	}{
		{
			name: "merge multiple uint16 slices",
			slices: [][]uint16{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMergeUint32s(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]uint32
		expected []uint32
	}{
		{
			name: "merge multiple uint32 slices",
			slices: [][]uint32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMergeUint64s(t *testing.T) {
	testCases := []struct {
		name     string
		slices   [][]uint64
		expected []uint64
	}{
		{
			name: "merge multiple uint64 slices",
			slices: [][]uint64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge(tc.slices...)
			assert.Equal(t, tc.expected, result)
		})
	}
}
