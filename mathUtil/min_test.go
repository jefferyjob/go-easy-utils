package mathUtil

import "testing"

func TestMinEmpty(t *testing.T) {
	s := []int{}
	var expected int = 0
	res := Min(s)
	if res != expected {
		t.Errorf("min error %d", res)
	}
}

func TestMin(t *testing.T) {
	s := []int{3, 7, 1, 9, 3, 0, 2, 2}
	var expected int = 0
	res := Min(s)
	if res != expected {
		t.Errorf("min error %d", res)
	}
}
