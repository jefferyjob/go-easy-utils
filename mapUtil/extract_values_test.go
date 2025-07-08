package mapUtil

import (
	"sort"
	"testing"
)

func TestExtractValues(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []int
	}{
		{
			name:     "Normal map",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty map",
			input:    map[string]int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    map[string]int{"x": 10},
			expected: []int{10},
		},
		{
			name:     "Duplicate values",
			input:    map[string]int{"a": 5, "b": 5, "c": 5},
			expected: []int{5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractValues(tt.input)

			// 对结果排序，因为 map 迭代顺序是随机的
			sort.Ints(result)
			sort.Ints(tt.expected)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d values, got %d", len(tt.expected), len(result))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestExtractValues2(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		expected []string
	}{
		{
			name:     "int和string类型的map",
			input:    map[int]string{1: "one", 2: "two"},
			expected: []string{"one", "two"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractValues(tt.input)

			// 对结果排序，因为 map 迭代顺序是随机的
			sort.Strings(result)
			sort.Strings(tt.expected)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d values, got %d", len(tt.expected), len(result))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("At index %d: expected %s, got %s", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestExtractValues3(t *testing.T) {
	type Point struct{ X, Y int }

	tests := []struct {
		name     string
		input    map[string]Point
		expected []Point
	}{
		{
			name: "结构体类型的map",
			input: map[string]Point{
				"origin": {X: 0, Y: 0},
				"top":    {X: 0, Y: 10},
			},
			expected: []Point{
				{X: 0, Y: 0},
				{X: 0, Y: 10},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractValues(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
