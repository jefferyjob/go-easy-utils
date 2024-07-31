package anyUtil

import "testing"

func TestAnyToStr(t *testing.T) {
	iPtr := 90
	testCases := []struct {
		name     string
		input    any
		expected string
	}{
		{"String", "hello", "hello"},
		{"Integer", 123, "123"},
		{"Int8", int8(8), "8"},
		{"Int16", int16(16), "16"},
		{"Int32", int32(32), "32"},
		{"Int64", int64(64), "64"},
		{"Uint", uint(10), "10"},
		{"Uint8", uint8(80), "80"},
		{"Uint16", uint16(160), "160"},
		{"Uint32", uint32(320), "320"},
		{"Uint64", uint64(640), "640"},
		{"Float32", float32(3.14159), "3.14159"},
		{"Float64", float64(2.71828), "2.71828"},
		{"Complex64", complex64(1 + 2i), "(1+2i)"},
		{"Complex128", complex128(3 + 4i), "(3+4i)"},
		{"NilPointer", (*int)(nil), ""},
		{"PointerToInt", &iPtr, "90"},
		{"Boolean", true, "true"},
		{"Nil", nil, ""},
		{"Channel", make(chan int), ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := AnyToStr(tc.input)
			if result != tc.expected {
				t.Errorf("Unexpected result. Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
