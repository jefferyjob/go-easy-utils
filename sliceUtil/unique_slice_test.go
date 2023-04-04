package sliceUtil

import (
	"reflect"
	"testing"
)

func TestUniqueSlice(t *testing.T) {
	// 测试整型切片去重
	s1 := []int{1, 2, 3, 2, 4, 4, 5}
	s1 = UniqueSlice(s1)
	if !reflect.DeepEqual(s1, []int{1, 2, 3, 4, 5}) {
		t.Errorf("UniqueSlice(%v) = %v, expected %v", s1, s1, []int{1, 2, 3, 4, 5})
	}

	// 测试字符串切片去重
	s2 := []string{"a", "b", "c", "b", "d", "d", "e"}
	s2 = UniqueSlice(s2)
	if !reflect.DeepEqual(s2, []string{"a", "b", "c", "d", "e"}) {
		t.Errorf("UniqueSlice(%v) = %v, expected %v", s2, s2, []string{"a", "b", "c", "d", "e"})
	}
}

func TestUniqueSliceInt(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 1, 1}, []int{1}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		got := UniqueSlice(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("UniqueSliceInt(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
