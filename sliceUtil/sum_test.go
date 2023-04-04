package sliceUtil

import (
	"testing"
)

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
			result := Sum(tc.input)
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
			result := Sum(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, result)
			}
		})
	}
}

// 测试SumInt16Slice函数
func TestSumInt16Slice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int16
		want  int16
	}{
		{"empty slice", []int16{}, 0},
		{"one element slice", []int16{1}, 1},
		{"multiple element slice", []int16{1, 2, 3}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.slice)
			if got != tt.want {
				t.Errorf("SumInt16Slice(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}

// 测试SumInt32Slice函数
func TestSumInt32Slice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int32
		want  int32
	}{
		{"empty slice", []int32{}, 0},
		{"one element slice", []int32{1}, 1},
		{"multiple element slice", []int32{1, 2, 3}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.slice)
			if got != tt.want {
				t.Errorf("SumInt32Slice(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}

// 测试SumInt64Slice函数
func TestSumInt64Slice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int64
		want  int64
	}{
		{"empty slice", []int64{}, 0},
		{"one element slice", []int64{1}, 1},
		{"multiple element slice", []int64{1, 2, 3}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.slice)
			if got != tt.want {
				t.Errorf("SumInt64Slice(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}

// 测试SumFloat32Slice函数
func TestSumFloat32Slice(t *testing.T) {
	tests := []struct {
		name  string
		slice []float32
		want  float32
	}{
		{"empty slice", []float32{}, 0},
		{"one element slice", []float32{1.0}, 1.0},
		{"multiple element slice", []float32{1.0, 2.0, 3.0}, 6.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.slice)
			if got != tt.want {
				t.Errorf("SumFloat32Slice(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}

// 测试SumFloat64Slice函数
func TestSumFloat64Slice(t *testing.T) {
	tests := []struct {
		name  string
		slice []float64
		want  float64
	}{
		{"empty slice", []float64{}, 0},
		{"one element slice", []float64{1.0}, 1.0},
		{"multiple element slice", []float64{1.0, 2.0, 3.0}, 6.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.slice)
			if got != tt.want {
				t.Errorf("SumFloat64Slice(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}
