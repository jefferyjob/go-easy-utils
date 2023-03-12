package sliceUtil

import (
	"reflect"
	"testing"
)

// 测试MergeSlices
func TestMergeSlices(t *testing.T) {
	slice1 := []interface{}{1, 2, 3}
	slice2 := []interface{}{4, 5, 6}
	slice3 := []interface{}{7, 8, 9}
	expected := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := MergeSlices(slice1, slice2, slice3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// 测试MergeStrSlices
func TestMergeStrSlices(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"d", "e", "f"}
	slice3 := []string{"g", "h", "i"}
	expected := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	result := MergeStrSlices(slice1, slice2, slice3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
