package sliceUtil

import "testing"

func TestSumIntSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Empty Slice",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "One Element Slice",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "Multiple Elements Slice",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SumIntSlice(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, result)
			}
		})
	}
}

func TestSumInt8Slice(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int8
		expected int8
	}{
		{
			name:     "Empty Slice",
			input:    []int8{},
			expected: 0,
		},
		{
			name:     "One Element Slice",
			input:    []int8{1},
			expected: 1,
		},
		{
			name:     "Multiple Elements Slice",
			input:    []int8{1, 2, 3, 4, 5},
			expected: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SumInt8Slice(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, result)
			}
		})
	}
}
