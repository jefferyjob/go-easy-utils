package strUtil

import "testing"

func TestStrToUint(t *testing.T) {
	tests := []struct {
		input    string
		expected uint
	}{
		{"123", 123},
		{"0", 0},
		{"-1", 0},
	}

	for _, test := range tests {
		result := StrToUint(test.input)
		if result != test.expected {
			t.Errorf("StrToUint(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestStrToUint8(t *testing.T) {
	tests := []struct {
		input    string
		expected uint8
	}{
		{"123", 123},
		{"0", 0},
		{"-1", 0},
		{"256", 0},
	}

	for _, test := range tests {
		result := StrToUint8(test.input)
		if result != test.expected {
			t.Errorf("StrToUint8(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestStrToUint16(t *testing.T) {
	tests := []struct {
		input    string
		expected uint16
	}{
		{"12345", 12345},
		{"0", 0},
		{"-1", 0},
		{"65536", 0},
	}

	for _, test := range tests {
		result := StrToUint16(test.input)
		if result != test.expected {
			t.Errorf("StrToUint16(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestStrToUint32(t *testing.T) {
	tests := []struct {
		input    string
		expected uint32
	}{
		{"1234567", 1234567},
		{"0", 0},
		{"-1", 0},
		{"4294967296", 0},
	}

	for _, test := range tests {
		result := StrToUint32(test.input)
		if result != test.expected {
			t.Errorf("StrToUint32(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestStrToUint64(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput uint64
	}{
		{"12345", 12345},
		{"0", 0},
		{"18446744073709551615", 18446744073709551615},
		{"-1", 0},
		{"not a number", 0},
	}

	for _, tc := range testCases {
		output := StrToUint64(tc.input)
		if output != tc.expectedOutput {
			t.Errorf("StrToUint64(%q) = %d; want %d", tc.input, output, tc.expectedOutput)
		}
	}
}
