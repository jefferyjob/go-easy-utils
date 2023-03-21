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

func TestInIntSlices(t *testing.T) {
	slices := []int{1, 2, 3, 4, 5}
	if !InIntSlices(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InIntSlices(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInInt8Slices(t *testing.T) {
	slices := []int8{1, 2, 3, 4, 5}
	if !InInt8Slices(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InInt8Slices(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInInt16Slices(t *testing.T) {
	slices := []int16{1, 2, 3, 4, 5}
	if !InInt16Slices(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InInt16Slices(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInInt32Slices(t *testing.T) {
	slices := []int32{1, 2, 3, 4, 5}
	if !InInt32Slices(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InInt32Slices(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInInt64Slices(t *testing.T) {
	slices := []int64{1, 2, 3, 4, 5}
	if !InInt64Slices(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InInt64Slices(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInUintSlices(t *testing.T) {
	slices := []uint{1, 2, 3, 4, 5}
	if !InUintSlices(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InUintSlices(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInUint8Slices(t *testing.T) {
	slices := []uint8{1, 2, 3, 4, 5}
	if !InUint8Slices(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InUint8Slices(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInUint16Slices(t *testing.T) {
	slices := []uint16{1, 2, 3, 4, 5}
	if !InUint16Slices(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InUint16Slices(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInUint32Slices(t *testing.T) {
	slices := []uint32{1, 2, 3, 4, 5}
	if !InUint32Slices(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InUint32Slices(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}

func TestInUint64Slices(t *testing.T) {
	slices := []uint64{1, 2, 3, 4, 5}
	if !InUint64Slices(3, slices) {
		t.Errorf("Expected true, but got false")
	}

	if InUint64Slices(6, slices) {
		t.Errorf("Expected false, but got true")
	}
}
