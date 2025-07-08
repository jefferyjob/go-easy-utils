package mapx

import (
	"sort"
	"testing"
)

func TestKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{
			name:     "Normal map",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Empty map",
			input:    map[string]int{},
			expected: []string{},
		},
		{
			name:     "Single element",
			input:    map[string]int{"x": 10},
			expected: []string{"x"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Keys(tt.input)

			// 对结果排序，因为 map 迭代顺序是随机的
			sort.Strings(result)
			sort.Strings(tt.expected)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d keys, got %d", len(tt.expected), len(result))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %q, got %q", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestKeys2(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		expected []int
	}{
		{
			name:     "int和string类型的map",
			input:    map[int]string{1: "one", 2: "two"},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Keys(tt.input)

			// 对结果排序，因为 map 迭代顺序是随机的
			sort.Ints(result)
			sort.Ints(tt.expected)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d keys, got %d", len(tt.expected), len(result))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %q, got %q", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestKeys3(t *testing.T) {
	tests := []struct {
		name     string
		input    map[float64]bool
		expected []float64
	}{
		{
			name:     "float64和bool类型的map",
			input:    map[float64]bool{3.14: true, 2.71: false},
			expected: []float64{2.71, 3.14},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Keys(tt.input)

			// 对结果排序，因为 map 迭代顺序是随机的
			sort.Float64s(result)
			sort.Float64s(tt.expected)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d keys, got %d", len(tt.expected), len(result))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %f, got %f", i, tt.expected[i], result[i])
				}
			}
		})
	}
}
