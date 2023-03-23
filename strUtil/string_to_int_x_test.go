package strUtil

import (
	"bytes"
	"testing"
)

func TestStrToInt(t *testing.T) {
	expected := 123
	result := StrToInt("123")
	if result != expected {
		t.Errorf("StrToInt test failed, expected %d but got %d", expected, result)
	}
}

func TestStrToInt8(t *testing.T) {
	expected := int8(123)
	result := StrToInt8("123")
	if result != expected {
		t.Errorf("StrToInt8 test failed, expected %d but got %d", expected, result)
	}
}

func TestStrToInt16(t *testing.T) {
	expected := int16(123)
	result := StrToInt16("123")
	if result != expected {
		t.Errorf("StrToInt16 test failed, expected %d but got %d", expected, result)
	}
}

func TestStrToInt32(t *testing.T) {
	expected := int32(123)
	result := StrToInt32("123")
	if result != expected {
		t.Errorf("StrToInt32 test failed, expected %d but got %d", expected, result)
	}
}

func TestStrToInt64(t *testing.T) {
	expected := int64(123)
	result := StrToInt64("123")
	if result != expected {
		t.Errorf("StrToInt64 test failed, expected %d but got %d", expected, result)
	}
}

func TestStrToBytes(t *testing.T) {
	expected := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	result := StrToBytes("hello")
	if !bytes.Equal(result, expected) {
		t.Errorf("StrToBytes test failed, expected %v but got %v", expected, result)
	}
}

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
