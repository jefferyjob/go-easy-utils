package mathUtil

import (
	"testing"
)

func TestCeilFloat32(t *testing.T) {
	var input float32 = -12.4
	var expected int = -12
	res := Ceil(input)
	if res != expected {
		t.Errorf("ceil error")
	}
}

func TestCeilFloat64(t *testing.T) {
	var input float64 = 12.4
	var expected int = 13
	res := Ceil(input)
	if res != expected {
		t.Errorf("ceil error")
	}
}
