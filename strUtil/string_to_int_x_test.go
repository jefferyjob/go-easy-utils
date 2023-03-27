package strUtil

import (
	"testing"
)

func TestStrToInt(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected int64
	}{
		{"0", 0},
		{"1", 1},
		{"-1", -1},
		{"invalid", 0},
	}

	// Test loop
	for _, tc := range testCases {
		// Call the function
		got := StrToInt64(tc.input)

		// Check the result
		if got != tc.expected {
			t.Errorf("Unexpected result: StrToInt(%v) = %v, expected %v", tc.input, got, tc.expected)
		}
	}
}

func TestStrToInt8(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected int64
	}{
		{"0", 0},
		{"1", 1},
		{"-1", -1},
		{"invalid", 0},
	}

	// Test loop
	for _, tc := range testCases {
		// Call the function
		got := StrToInt64(tc.input)

		// Check the result
		if got != tc.expected {
			t.Errorf("Unexpected result: StrToInt8(%v) = %v, expected %v", tc.input, got, tc.expected)
		}
	}
}

func TestStrToInt16(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected int64
	}{
		{"0", 0},
		{"1", 1},
		{"-1", -1},
		{"invalid", 0},
	}

	// Test loop
	for _, tc := range testCases {
		// Call the function
		got := StrToInt64(tc.input)

		// Check the result
		if got != tc.expected {
			t.Errorf("Unexpected result: StrToInt16(%v) = %v, expected %v", tc.input, got, tc.expected)
		}
	}
}

func TestStrToInt32(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected int64
	}{
		{"0", 0},
		{"1", 1},
		{"-1", -1},
		{"invalid", 0},
	}

	// Test loop
	for _, tc := range testCases {
		// Call the function
		got := StrToInt64(tc.input)

		// Check the result
		if got != tc.expected {
			t.Errorf("Unexpected result: StrToInt32(%v) = %v, expected %v", tc.input, got, tc.expected)
		}
	}
}

func TestStrToInt64(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected int64
	}{
		{"0", 0},
		{"1", 1},
		{"-1", -1},
		{"9223372036854775807", 9223372036854775807},
		{"-9223372036854775808", -9223372036854775808},
		{"2147483647", 2147483647},
		{"-2147483648", -2147483648},
		{"invalid", 0},
		{"9223372036854775808", 0},
		{"-9223372036854775809", 0},
	}

	// Test loop
	for _, tc := range testCases {
		// Call the function
		got := StrToInt64(tc.input)

		// Check the result
		if got != tc.expected {
			t.Errorf("Unexpected result: StrToInt64(%v) = %v, expected %v", tc.input, got, tc.expected)
		}
	}
}
