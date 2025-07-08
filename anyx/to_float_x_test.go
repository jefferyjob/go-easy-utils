package anyx

import (
	"errors"
	"math"
	"testing"
)

func TestToFloat32(t *testing.T) {
	testCases := []struct {
		name  string
		input any
		want  float32
		err   error
	}{
		{"nil", nil, 0, nil},
		{"float32", float32(3.14), 3.14, nil},
		{"float64", float64(6.28), 6.28, nil},
		{"int", int(42), 42, nil},
		{"int8", int8(8), 8, nil},
		{"int16", int16(16), 16, nil},
		{"int32", int32(32), 32, nil},
		{"int64", int64(64), 64, nil},
		{"uint", uint(24), 24, nil},
		{"uint8", uint8(8), 8, nil},
		{"uint16", uint16(16), 16, nil},
		{"uint32", uint32(32), 32, nil},
		{"uint64", uint64(64), 64, nil},
		{"string", "3.14", 3.14, nil},
		{"math.MaxFloat64", math.MaxFloat64, 0, ErrValOut},
		{"channel", make(chan int), 0, ErrType},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := ToFloat32(test.input)
			if got != test.want {
				t.Errorf("ToFloat32(%v) = %v; want %v", test.input, got, test.want)
			}
			if !errors.Is(err, test.err) {
				if err != nil && test.err != nil {
					if err.Error() != test.err.Error() {
						t.Errorf("ToFloat32(%v) error = %v; want %v", test.input, err, test.err)
					}
				} else {
					t.Errorf("ToFloat32(%v) error = %v; want %v", test.input, err, test.err)
				}
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	iPtr := 90

	tests := []struct {
		name  string
		input any
		want  float64
		err   error
	}{
		{"nil", nil, 0, nil},
		{"float64", float64(6.28), 6.28, nil},
		{"int", int(42), 42, nil},
		{"int8", int8(8), 8, nil},
		{"int16", int16(16), 16, nil},
		{"int32", int32(32), 32, nil},
		{"int64", int64(64), 64, nil},
		{"uint", uint(24), 24, nil},
		{"uint8", uint8(8), 8, nil},
		{"uint16", uint16(16), 16, nil},
		{"uint32", uint32(32), 32, nil},
		{"uint64", uint64(64), 64, nil},
		{"string", "3.14", 3.14, nil},
		{"true", true, 1, nil},
		{"false", false, 0, nil},
		{"complex64", complex64(1 + 2i), 1, nil},
		{"complex128", complex128(1 + 2i), 1, nil},
		{"pointer", &iPtr, 90, nil},
		{"nil pointer", (*int)(nil), 0, nil},
		{"invalid string", "abc", 0, ErrSyntax},
		{"slice", []int{}, 0, ErrType},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ToFloat64(test.input)
			if got != test.want {
				t.Errorf("ToFloat64(%v) = %v; want %v", test.input, got, test.want)
			}
			if err != test.err {
				if err != nil && test.err != nil {
					if err.Error() != test.err.Error() {
						t.Errorf("ToFloat64(%v) error = %v; want %v", test.input, err, test.err)
					}
				} else {
					t.Errorf("ToFloat64(%v) error = %v; want %v", test.input, err, test.err)
				}
			}
		})
	}
}
