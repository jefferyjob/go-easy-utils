package mathUtil

import "testing"

func TestRand(t *testing.T) {
	minNum := 1
	maxNum := 100
	for i := 0; i < 1000; i++ {
		num := Rand(minNum, maxNum)
		if num < minNum || num > maxNum {
			t.Errorf("randRange() returned %d, which is outside the range of [%d, %d]", num, minNum, maxNum)
		}
	}
}
