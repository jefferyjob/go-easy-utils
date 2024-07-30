package mathUtil

import (
	"testing"
)

func TestAbsInt64(t *testing.T) {
	var input int64 = -12
	var expected int64 = 12
	res := Abs(input)
	if res != expected {
		t.Errorf("abs error")
	}
}

func TestAbsFloat32(t *testing.T) {
	var input float32 = -12.4
	var expected float32 = 12.4
	res := Abs(input)
	if res != expected {
		t.Errorf("abs error")
	}
}
