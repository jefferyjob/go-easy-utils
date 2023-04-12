package mathUtil

import "testing"

func TestFloorFloat32(t *testing.T) {
	var input float32 = -12.4
	var expected int = -13
	res := Floor(input)
	if res != expected {
		t.Errorf("floor error %d", res)
	}
}

func TestFloorFloat64(t *testing.T) {
	var input float64 = 12.4
	var expected int = 12
	res := Floor(input)
	if res != expected {
		t.Errorf("floor error %d", res)
	}
}
