package slicex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueString(t *testing.T) {
	testCases := []struct {
		name string
		in   []string
		want []string
	}{
		{
			name: "字符串切片去重",
			in:   []string{"a", "b", "c", "b", "d", "d", "e"},
			want: []string{"a", "b", "c", "d", "e"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Unique(tc.in)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestUniqueInt(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "重复元素", in: []int{1, 2, 3, 2, 1}, want: []int{1, 2, 3}},
		{name: "所有元素相同", in: []int{1, 1, 1, 1}, want: []int{1}},
		{name: "无重复元素", in: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
		{name: "空切片", in: []int{}, want: []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Unique(tc.in)
			assert.Equal(t, tc.want, res)
		})
	}
}
