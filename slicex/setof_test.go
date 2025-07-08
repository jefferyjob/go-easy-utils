package slicex

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	testCases := []struct {
		name   string
		s      []int
		slices [][]int
		want   []int
	}{
		{
			name:   "第一个切片与空切片列表的比较",
			s:      []int{1, 2, 3},
			slices: [][]int{},
			want:   []int{1, 2, 3},
		},
		{
			name:   "第一个切片与包含部分相同元素的切片列表的比较",
			s:      []int{1, 2, 3, 4, 5},
			slices: [][]int{{4, 5}, {6, 7}},
			want:   []int{1, 2, 3},
		},
		{
			name:   "第一个切片与包含全部相同元素的切片列表的比较",
			s:      []int{1, 2, 3},
			slices: [][]int{{1, 2, 3}},
			want:   []int{},
		},
		{
			name:   "第一个切片与不包含任何相同元素的切片列表的比较",
			s:      []int{1, 2, 3},
			slices: [][]int{{4, 5, 6}},
			want:   []int{1, 2, 3},
		},
		{
			name:   "空切片与任何切片列表的比较",
			s:      []int{},
			slices: [][]int{{1, 2, 3}},
			want:   []int{},
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := Diff(tc.s, tc.slices...)
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Diff(%v, %v) = %v, want %v", tc.s, tc.slices, got, tc.want)
			}
		})
	}
}

func TestSymmetricDiff(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			name:     "两个切片计算对称差集",
			slices:   [][]int{{1, 2, 3}, {2, 3}},
			expected: []int{1},
		},
		{
			name:     "多个切片计算对称差集",
			slices:   [][]int{{2, 4}, {3}, {1, 2, 3, 4, 5}},
			expected: []int{1, 5},
		},
		{
			name:     "No Symmetric Difference",
			slices:   [][]int{{1, 2, 3}, {1, 2, 3}},
			expected: []int{},
		},
		{
			name:     "Empty Base ",
			slices:   [][]int{{}, {1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty Comparison s",
			slices:   [][]int{{1, 2, 3}, {}, {}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "All Empty s",
			slices:   [][]int{{}, {}, {}},
			expected: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SymmetricDiff(tt.slices...)
			if len(got) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Diff() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			name:     "Single Intersection",
			slices:   [][]int{{1, 2, 3}, {2, 3, 4}},
			expected: []int{2, 3},
		},
		{
			name:     "Multiple Intersections",
			slices:   [][]int{{2, 3, 4}, {2, 3, 5}, {1, 2, 2, 3}},
			expected: []int{2, 3},
		},
		{
			name:     "No Intersection",
			slices:   [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: []int{},
		},
		{
			name:     "Empty Base ",
			slices:   [][]int{{}, {1, 2, 3}},
			expected: []int{},
		},
		{
			name:     "Empty Comparison s",
			slices:   [][]int{{1, 2, 3}, {}, {}},
			expected: []int{},
		},
		{
			name:     "All Empty s",
			slices:   [][]int{{}, {}, {}},
			expected: []int{},
		},
		{
			name:     "All Same Elements",
			slices:   [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
			expected: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersect(tt.slices...)
			if len(got) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Intersect() = %v, want %v", got, tt.expected)
			}
		})
	}
}
