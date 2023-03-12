package floatUtil

import (
	"math"
	"testing"
)

func TestFloat64Round(t *testing.T) {
	tests := []struct {
		input     float64
		precision int
		expected  float64
	}{
		{2.3456, 2, 2.35},
		{0.1234, 1, 0.1},
		{1234.5678, 0, 1235},
	}

	for _, test := range tests {
		result := Float64Round(test.input, test.precision)
		if math.Abs(result-test.expected) > 1e-6 {
			t.Errorf("Float64Round(%f,%d)=%f; expected %f", test.input, test.precision, result, test.expected)
		}
	}
}

func TestFloat64Ceil(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{2.3456, 2.35},
		{0.1234, 0.13},
		{1234.5678, 1234.57},
	}

	for _, test := range tests {
		result := Float64Ceil(test.input)
		if math.Abs(result-test.expected) > 1e-6 {
			t.Errorf("Float64Ceil(%f)=%f; expected %f", test.input, result, test.expected)
		}
	}
}
