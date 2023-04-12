package mathUtil

import "testing"

func TestRoundFloat32(t *testing.T) {
	var input float32 = -12.4
	var expected int = -12
	res := Round(input)
	if res != expected {
		t.Errorf("Round error %d", res)
	}
}

func TestRoundFloat64(t *testing.T) {
	var input float64 = 12.5
	var expected int = 13
	res := Round(input)
	if res != expected {
		t.Errorf("Round error %d", res)
	}
}
