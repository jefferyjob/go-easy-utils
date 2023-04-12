package mathUtil

import "testing"

func TestRand(t *testing.T) {
	min := 1
	max := 100
	for i := 0; i < 1000; i++ {
		num := Rand(min, max)
		if num < min || num > max {
			t.Errorf("randRange() returned %d, which is outside the range of [%d, %d]", num, min, max)
		}
	}
}
