package sliceUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunk(t *testing.T) {
	testCases := []struct {
		name  string
		slice []any
		size  int
		want  [][]any
	}{
		{
			name:  "切片小于块大小",
			slice: []any{1, 2},
			size:  3,
			want:  [][]any{{1, 2}},
		},
		{
			name:  "切片等于块大小",
			slice: []any{1, 2, 3, 4, 5, 6},
			size:  3,
			want:  [][]any{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name:  "切片大于块大小",
			slice: []any{1, 2, 3, 4, 5, 6, 7},
			size:  3,
			want:  [][]any{{1, 2, 3}, {4, 5, 6}, {7}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestChunkStr(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []string
		size     int
		expected [][]string
	}{
		{
			name:     "切片大小为3",
			slice:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"},
			size:     3,
			expected: [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestChunkInt(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int
		size     int
		expected [][]int
	}{
		{
			name:     "切片大小为4",
			slice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			size:     4,
			expected: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestChunkInt8(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int8
		size     int
		expected [][]int8
	}{
		{
			name:     "切片大小为4",
			slice:    []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			size:     4,
			expected: [][]int8{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestChunkInt32(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int32
		size     int
		expected [][]int32
	}{
		{
			name:     "切片大小为4",
			slice:    []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			size:     4,
			expected: [][]int32{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestChunkInt64(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int64
		size     int
		expected [][]int64
	}{
		{
			name:     "切片大小为4",
			slice:    []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			size:     4,
			expected: [][]int64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestChunkUint(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []uint
		size     int
		expected [][]uint
	}{
		{
			name:     "切片大小为4",
			slice:    []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			size:     4,
			expected: [][]uint{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestChunkUint8(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []uint8
		size     int
		expected [][]uint8
	}{
		{
			name:     "切片大小为4",
			slice:    []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			size:     4,
			expected: [][]uint8{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestChunkUint16(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []uint16
		size     int
		expected [][]uint16
	}{
		{
			name:     "切片大小为4",
			slice:    []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			size:     4,
			expected: [][]uint16{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestChunkUint32(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []uint32
		size     int
		expected [][]uint32
	}{
		{
			name:     "切片大小为4",
			slice:    []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			size:     4,
			expected: [][]uint32{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestChunkUint64(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []uint64
		size     int
		expected [][]uint64
	}{
		{
			name:     "切片大小为4",
			slice:    []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			size:     4,
			expected: [][]uint64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChunkSlice(tc.slice, tc.size)
			assert.Equal(t, tc.expected, res)
		})
	}
}
