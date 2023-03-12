package sliceUtil

import (
	"testing"
)

func TestInSlices(t *testing.T) {
	slices := []interface{}{1, "hello", 3.14}
	if !InSlices(1, slices) {
		t.Errorf("1 should be in slices %v", slices)
	}
	if !InSlices("hello", slices) {
		t.Errorf("hello should be in slices %v", slices)
	}
	if InSlices(2, slices) {
		t.Errorf("2 should not be in slices %v", slices)
	}
}

func TestInStrSlices(t *testing.T) {
	slices := []string{"apple", "banana", "cherry"}
	if !InStrSlices("apple", slices) {
		t.Errorf("apple should be in slices %v", slices)
	}
	if !InStrSlices("banana", slices) {
		t.Errorf("banana should be in slices %v", slices)
	}
	if InStrSlices("orange", slices) {
		t.Errorf("orange should not be in slices %v", slices)
	}
}
