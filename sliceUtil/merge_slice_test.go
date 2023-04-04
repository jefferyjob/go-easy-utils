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
	result := MergeSlice(slice1, slice2, slice3)
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
	result := MergeSlice(slice1, slice2, slice3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMergeIntSlices(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := []int{7, 8, 9}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MergeSlice(slice1, slice2, slice3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMergeInt8Slices(t *testing.T) {
	slice1 := []int8{1, 2, 3}
	slice2 := []int8{4, 5, 6}
	slice3 := []int8{7, 8, 9}
	expected := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MergeSlice(slice1, slice2, slice3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMergeInt16Slices(t *testing.T) {
	slice1 := []int16{1, 2, 3}
	slice2 := []int16{4, 5, 6}
	slice3 := []int16{7, 8, 9}
	expected := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MergeSlice(slice1, slice2, slice3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMergeInt32Slices(t *testing.T) {
	slice1 := []int32{1, 2, 3}
	slice2 := []int32{4, 5, 6}
	slice3 := []int32{7, 8, 9}
	expected := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MergeSlice(slice1, slice2, slice3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMergeInt64Slices(t *testing.T) {
	slice1 := []int64{1, 2, 3}
	slice2 := []int64{4, 5, 6}
	slice3 := []int64{7, 8, 9}
	expected := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MergeSlice(slice1, slice2, slice3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMergeUintSlices(t *testing.T) {
	slice1 := []uint{1, 2, 3}
	slice2 := []uint{4, 5, 6}
	slice3 := []uint{7, 8, 9}
	expected := []uint{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MergeSlice(slice1, slice2, slice3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMergeUint8Slices(t *testing.T) {
	slice1 := []uint8{1, 2, 3}
	slice2 := []uint8{4, 5, 6}
	slice3 := []uint8{7, 8, 9}
	expected := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MergeSlice(slice1, slice2, slice3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMergeUint16Slices(t *testing.T) {
	slice1 := []uint16{1, 2, 3}
	slice2 := []uint16{4, 5, 6}
	slice3 := []uint16{7, 8, 9}
	expected := []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MergeSlice(slice1, slice2, slice3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMergeUint32Slices(t *testing.T) {
	slice1 := []uint32{1, 2, 3}
	slice2 := []uint32{4, 5, 6}
	slice3 := []uint32{7, 8, 9}
	expected := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MergeSlice(slice1, slice2, slice3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMergeUint64Slices(t *testing.T) {
	slice1 := []uint64{1, 2, 3}
	slice2 := []uint64{4, 5, 6}
	slice3 := []uint64{7, 8, 9}
	expected := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := MergeSlice(slice1, slice2, slice3)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
