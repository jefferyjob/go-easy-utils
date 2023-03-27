package anyUtil

import "testing"

func TestAnyToStr(t *testing.T) {
	iPtr := 90
	testCases := []struct {
		input    interface{}
		expected string
	}{
		{"hello", "hello"},
		{123, "123"},
		{int8(8), "8"},
		{int16(16), "16"},
		{int32(32), "32"},
		{int64(64), "64"},
		{uint(10), "10"},
		{uint8(80), "80"},
		{uint16(160), "160"},
		{uint32(320), "320"},
		{uint64(640), "640"},
		{float32(3.14159), "3.14159"},
		{float64(2.71828), "2.71828"},
		{complex64(1 + 2i), "(1+2i)"},
		{complex128(3 + 4i), "(3+4i)"},
		{(*int)(nil), ""},
		{&iPtr, "90"},
		{true, "true"},
		{nil, ""},
	}

	for _, testCase := range testCases {
		result := AnyToStr(testCase.input)
		if result != testCase.expected {
			t.Errorf("Unexpected result. Expected: %v, but got: %v", testCase.expected, result)
		}
	}
}
