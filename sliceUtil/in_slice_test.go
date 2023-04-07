package sliceUtil

import (
	"testing"
)

func TestInStrSlices(t *testing.T) {
	slices := []string{"apple", "banana", "cherry"}
	if !InSlice("apple", slices) {
		t.Errorf("apple should be in slices %v", slices)
	}
	if !InSlice("banana", slices) {
		t.Errorf("banana should be in slices %v", slices)
	}
	if InSlice("orange", slices) {
		t.Errorf("orange should not be in slices %v", slices)
	}
}

func TestInIntSlices(t *testing.T) {
	slices := []int{1, 2, 3, 4, 5}
	if !InSlice(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InSlice(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInInt8Slices(t *testing.T) {
	slices := []int8{1, 2, 3, 4, 5}
	if !InSlice(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InSlice(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInInt16Slices(t *testing.T) {
	slices := []int16{1, 2, 3, 4, 5}
	if !InSlice(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InSlice(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInInt32Slices(t *testing.T) {
	slices := []int32{1, 2, 3, 4, 5}
	if !InSlice(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InSlice(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInInt64Slices(t *testing.T) {
	slices := []int64{1, 2, 3, 4, 5}
	if !InSlice(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InSlice(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInUintSlices(t *testing.T) {
	slices := []uint{1, 2, 3, 4, 5}
	if !InSlice(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InSlice(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInUint8Slices(t *testing.T) {
	slices := []uint8{1, 2, 3, 4, 5}
	if !InSlice(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InSlice(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInUint16Slices(t *testing.T) {
	slices := []uint16{1, 2, 3, 4, 5}
	if !InSlice(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InSlice(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInUint32Slices(t *testing.T) {
	slices := []uint32{1, 2, 3, 4, 5}
	if !InSlice(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InSlice(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInUint64Slices(t *testing.T) {
	slices := []uint64{1, 2, 3, 4, 5}
	if !InSlice(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InSlice(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}
