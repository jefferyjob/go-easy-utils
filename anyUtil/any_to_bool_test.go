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
