package randUtil

import (
	"math"
	"testing"
)

func TestRandomInt(t *testing.T) {
	min := 1
	max := 10

	for i := 0; i < 10000; i++ {
		result := RandomInt(min, max)

		if result < min || result > max {
			t.Errorf("RandomInt result out of range: %v", result)
		}
	}
}

func TestRandomFloat(t *testing.T) {
	min := 1.0
	max := 10.0

	for i := 0; i < 10000; i++ {
		result := RandomFloat(min, max)

		if result < min || result > max {
			t.Errorf("RandomFloat result out of range: %v", result)
		}

		if math.IsNaN(result) {
			t.Errorf("RandomFloat result is NaN")
		}
	}
}
