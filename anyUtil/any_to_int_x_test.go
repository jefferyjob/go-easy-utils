package anyUtil

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println(AnyToInt("9223372036854775806"))
}

func TestAnyToInt(t *testing.T) {
	type testCase struct {
		input    interface{}
		expected int
		err      error
	}

	testCases := []testCase{
		{input: 123, expected: 123, err: nil},
		{input: int8(123), expected: 123, err: nil},
		{input: int16(123), expected: 123, err: nil},
		{input: int32(123), expected: 123, err: nil},
		{input: int64(123), expected: 123, err: nil},
		{input: uint(123), expected: 123, err: nil},
		{input: uint8(123), expected: 123, err: nil},
		{input: uint16(123), expected: 123, err: nil},
		{input: uint32(123), expected: 123, err: nil},
		{input: uint64(123), expected: 123, err: nil},
		{input: float32(123.45), expected: 123, err: nil},
		{input: float64(123.45), expected: 123, err: nil},
		{input: "123", expected: 123, err: nil},
		{input: "abc", expected: 0, err: ErrSyntax},
	}

	for _, tc := range testCases {
		actual, err := AnyToInt(tc.input)

		if err != nil && tc.err == nil {
			t.Errorf("AnyToInt(%v) returned unexpected error: %v", tc.input, err)
			continue
		}

		if err == nil && tc.err != nil {
			t.Errorf("AnyToInt(%v) expected error but got nil", tc.input)
			continue
		}

		if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("AnyToInt(%v) returned wrong error. Expected %v, but got %v", tc.input, tc.err, err)
			continue
		}

		if actual != tc.expected {
			t.Errorf("AnyToInt(%v) returned %v, expected %v", tc.input, actual, tc.expected)
		}
	}
}

func TestAnyToInt8(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		input    interface{}
		expected int8
		err      error
	}{
		{
			name:     "positive integer",
			input:    42,
			expected: 42,
			err:      nil,
		},
		{
			name:     "negative integer",
			input:    -42,
			expected: -42,
			err:      nil,
		},
		{
			name:     "positive integer out of range",
			input:    128,
			expected: 0,
			err:      ErrValOut,
		},
		{
			name:     "negative integer out of range",
			input:    -129,
			expected: 0,
			err:      ErrValOut,
		},
		{
			name:     "float",
			input:    3.14,
			expected: 3,
			err:      nil,
		},
		{
			name:     "string",
			input:    "42",
			expected: 42,
			err:      nil,
		},
		{
			name:     "string out of range",
			input:    "128",
			expected: 0,
			err:      ErrValOut,
		},
		{
			name:     "invalid string",
			input:    "invalid",
			expected: 0,
			err:      ErrSyntax,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := AnyToInt8(tc.input)

			// Check error
			if err != nil {
				if !errors.Is(err, tc.err) {
					t.Errorf("Expected error: %v, but got: %v", tc.err, err)
				}
				//if err.Error() != tc.err.Error() {
				//	t.Errorf("Expected error: %v, but got: %v", tc.err, err)
				//}
			}

			// Check result
			if actual != tc.expected {
				t.Errorf("Expected result: %v, but got: %v", tc.expected, actual)
			}
		})
	}
}

func TestAnyToInt16(t *testing.T) {
	tests := []struct {
		input interface{}
		want  int16
		err   error
	}{
		{int(12345), 12345, nil},
		{int8(123), 123, nil},
		{int16(12345), 12345, nil},
		{int32(12345), 12345, nil},
		{int64(12345), 12345, nil},
		{uint(12345), 12345, nil},
		{uint8(123), 123, nil},
		{uint16(12345), 12345, nil},
		{uint32(12345), 12345, nil},
		{uint64(12345), 12345, nil},
		{float32(123.45), 123, nil},
		{float64(123.45), 123, nil},
		{"12345", 12345, nil},
		{"not a number", 0, strconv.ErrSyntax},
	}

	for _, tt := range tests {
		got, err := AnyToInt16(tt.input)

		if got != tt.want {
			t.Errorf("AnyToInt16(%v) = %v, want %v", tt.input, got, tt.want)
		}

		if !errors.Is(err, tt.err) {
			t.Errorf("AnyToInt16(%v) error = %v, want error %v", tt.input, err, tt.err)
		}
	}
}

func TestAnyToInt32(t *testing.T) {
	testCases := []struct {
		input         interface{}
		expectedValue int32
		expectedError error
	}{
		{int(123), 123, nil},
		{int64(2147483647), 2147483647, nil},
		{int64(2147483648), 0, ErrValOut},
		{int64(-2147483648), -2147483648, nil},
		{int64(-2147483649), 0, ErrValOut},
		{float64(123.45), 123, nil},
		{"123", 123, nil},
		{"-2147483649", 0, ErrValOut},
		{struct{}{}, 0, ErrType},
	}

	for _, testCase := range testCases {
		actualValue, actualError := AnyToInt32(testCase.input)
		if actualValue != testCase.expectedValue {
			t.Errorf("Unexpected value. Input: %v, Expected: %d, Actual: %d", testCase.input, testCase.expectedValue, actualValue)
		}
		if (actualError == nil && testCase.expectedError != nil) || (actualError != nil && testCase.expectedError == nil) || (actualError != nil && actualError.Error() != testCase.expectedError.Error()) {
			t.Errorf("Unexpected error. Input: %v, Expected: %v, Actual: %v", testCase.input, testCase.expectedError, actualError)
		}
	}
}

func TestAnyToInt64(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected int64
		err      bool
	}{
		// test cases for int values
		{int(42), int64(42), false},
		{int8(42), int64(42), false},
		{int16(42), int64(42), false},
		{int32(42), int64(42), false},
		{int64(42), int64(42), false},
		// test cases for uint values
		{uint(42), int64(42), false},
		{uint8(42), int64(42), false},
		{uint16(42), int64(42), false},
		{uint32(42), int64(42), false},
		{uint64(42), int64(42), false},
		// test cases for float values
		{float32(42.42), int64(42), false},
		{float64(42.42), int64(42), false},
		// test cases for string values
		{"42", int64(42), false},
		{"-42", int64(-42), false},
		{"42.42", int64(0), true}, // invalid syntax
		// unsupported type
		{struct{}{}, int64(0), true},
	}

	for _, testCase := range testCases {
		actual, err := AnyToInt64(testCase.input)
		if err != nil && !testCase.err {
			t.Errorf("unexpected error: %s", err)
		}
		if actual != testCase.expected {
			t.Errorf("expected %d but got %d", testCase.expected, actual)
		}
	}
}
