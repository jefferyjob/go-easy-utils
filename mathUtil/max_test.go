package mathUtil

import "testing"

func TestMax(t *testing.T) {
	s := []int{3, 7, 1, 9, 3, 0, 2, 2}
	var expected int = 9
	res := Max(s)
	if res != expected {
		t.Errorf("max error %d", res)
	}
}
