package anyUtil

import (
	"testing"
)

func TestToBool(t *testing.T) {
	iPtr := 90
	var tests = []struct {
		input any
		want  bool
	}{
		{true, true},
		{false, false},
		{int(-1), true},
		{int(1), true},
		{int(0), false},
		{int8(1), true},
		{int8(0), false},
		{int16(1), true},
		{int16(0), false},
		{int32(1), true},
		{int32(0), false},
		{int64(1), true},
		{int64(100), true},
		{int64(0), false},
		{uint(1), true},
		{uint(0), false},
		{uint8(1), true},
		{uint8(0), false},
		{uint16(1), true},
		{uint16(0), false},
		{uint32(1), true},
		{uint32(0), false},
		{uint64(1), true},
		{uint64(0), false},
		{float32(1.0), true},
		{float32(0.0), false},
		{float64(1.0), true},
		{float64(0.0), false},
		{"abc", true},
		{"true", true},
		{"false", false},
		{"", false},
		{nil, false},
		{complex64(1 + 1i), true},
		{complex64(0 + 0i), false},
		{complex128(1 + 1i), true},
		{complex128(0 + 0i), false},
		{(*int)(nil), false},
		{&iPtr, true},
		{[]int{}, false},
	}
	for _, test := range tests {
		if got := AnyToBool(test.input); got != test.want {
			t.Errorf("toBool(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestAnyToBool(t *testing.T) {
	iPtr := 90
	testCases := []struct {
		name  string
		input any
		want  bool
	}{
		{"bool true", true, true},
		{"bool false", false, false},
		{"int -1", int(-1), true},
		{"int 1", int(1), true},
		{"int 0", int(0), false},
		{"int8 1", int8(1), true},
		{"int8 0", int8(0), false},
		{"int16 1", int16(1), true},
		{"int16 0", int16(0), false},
		{"int32 1", int32(1), true},
		{"int32 0", int32(0), false},
		{"int64 1", int64(1), true},
		{"int64 100", int64(100), true},
		{"int64 0", int64(0), false},
		{"uint 1", uint(1), true},
		{"uint 0", uint(0), false},
		{"uint8 1", uint8(1), true},
		{"uint8 0", uint8(0), false},
		{"uint16 1", uint16(1), true},
		{"uint16 0", uint16(0), false},
		{"uint32 1", uint32(1), true},
		{"uint32 0", uint32(0), false},
		{"uint64 1", uint64(1), true},
		{"uint64 0", uint64(0), false},
		{"float32 1.0", float32(1.0), true},
		{"float32 0.0", float32(0.0), false},
		{"float64 1.0", float64(1.0), true},
		{"float64 0.0", float64(0.0), false},
		{"string abc", "abc", true},
		{"string true", "true", true},
		{"string false", "false", false},
		{"empty string", "", false},
		{"nil value", nil, false},
		{"complex64 1+1i", complex64(1 + 1i), true},
		{"complex64 0+0i", complex64(0 + 0i), false},
		{"complex128 1+1i", complex128(1 + 1i), true},
		{"complex128 0+0i", complex128(0 + 0i), false},
		{"nil pointer", (*int)(nil), false},
		{"non-nil pointer", &iPtr, true},
		{"empty slice", []int{}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := AnyToBool(tc.input); got != tc.want {
				t.Errorf("AnyToBool(%v) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}
