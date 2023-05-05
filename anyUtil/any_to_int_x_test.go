package anyUtil

import (
	"errors"
	"github.com/jefferyjob/go-easy-utils/v2"
	"math"
	"testing"
)

//func TestAnyToIntDemo(t *testing.T) {
//	fmt.Println(AnyToInt(math.MinInt32))
//}

func TestAnyToInt(t *testing.T) {
	type testCase struct {
		input    any
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
		{input: "abc", expected: 0, err: go_easy_utils.ErrSyntax},
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
		input    any
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
			err:      go_easy_utils.ErrValOut,
		},
		{
			name:     "negative integer out of range",
			input:    -129,
			expected: 0,
			err:      go_easy_utils.ErrValOut,
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
			err:      go_easy_utils.ErrValOut,
		},
		{
			name:     "invalid string",
			input:    "invalid",
			expected: 0,
			err:      go_easy_utils.ErrSyntax,
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
		input any
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
		{math.MinInt16 - 1, 0, go_easy_utils.ErrValOut},
		{"not a number", 0, go_easy_utils.ErrSyntax},
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
		input         any
		expectedValue int32
		expectedError error
	}{
		{int(123), 123, nil},
		{int64(2147483647), 2147483647, nil},
		{int64(2147483648), 0, go_easy_utils.ErrValOut},
		{int64(-2147483648), -2147483648, nil},
		{int64(-2147483649), 0, go_easy_utils.ErrValOut},
		{float64(123.45), 123, nil},
		{"123", 123, nil},
		{"-2147483649", 0, go_easy_utils.ErrValOut},
		{struct{}{}, 0, go_easy_utils.ErrType},
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
	iPtr := 90

	testCases := []struct {
		input    any
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
		// complex
		{complex64(1 + 2i), 1, false},
		{complex128(1 + 2i), 1, false},
		// point
		{(*int)(nil), 0, false},
		{nil, 0, false},
		{&iPtr, 90, false},
		// bool
		{true, 1, false},
		{false, 0, false},
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
