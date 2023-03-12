package intUtil

import "testing"

func TestIntToString(t *testing.T) {
	// Test cases with known inputs and expected outputs
	testCases := []struct {
		input    int
		expected string
	}{
		{0, "0"},
		{123, "123"},
		{-123, "-123"},
	}

	// Iterate through test cases
	for _, testCase := range testCases {
		// Call function with input and get result
		result := IntToString(testCase.input)

		// Compare result with expected output
		if result != testCase.expected {
			t.Errorf("IntToString(%d) = %s; expected %s", testCase.input, result, testCase.expected)
		}
	}
}

func TestInt8ToString(t *testing.T) {
	n := int8(42)
	expected := "42"
	result := Int8ToString(n)
	if result != expected {
		t.Errorf("Int8ToString(%d) = %s; expected %s", n, result, expected)
	}
}

func TestInt16ToString(t *testing.T) {
	n := int16(42)
	expected := "42"
	result := Int16ToString(n)
	if result != expected {
		t.Errorf("Int16ToString(%d) = %s; expected %s", n, result, expected)
	}
}

func TestInt32ToString(t *testing.T) {
	n := int32(42)
	expected := "42"
	result := Int32ToString(n)
	if result != expected {
		t.Errorf("Int32ToString(%d) = %s; expected %s", n, result, expected)
	}
}

func TestInt64ToString(t *testing.T) {
	n := int64(42)
	expected := "42"
	result := Int64ToString(n)
	if result != expected {
		t.Errorf("Int64ToString(%d) = %s; expected %s", n, result, expected)
	}
}
