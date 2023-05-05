package anyUtil

import (
	"github.com/jefferyjob/go-easy-utils/v2"
	"reflect"
	"testing"
)

func TestAnyToUint(t *testing.T) {
	// Test cases
	tests := []struct {
		input  any
		output uint
		err    error
	}{
		{10, 10, nil},
		{-5, 0, go_easy_utils.ErrUnsignedInt},
		{"20", 20, nil},
		{1.5, 1, nil},
		{make(chan int), 0, go_easy_utils.ErrType},
	}

	// Test loop
	for _, test := range tests {
		result, err := AnyToUint(test.input)
		if result != test.output || !reflect.DeepEqual(err, test.err) {
			t.Errorf("AnyToUint(%v) = (%v, %v), expected (%v, %v)",
				test.input, result, err, test.output, test.err)
		}
	}
}

func TestAnyToUint8(t *testing.T) {
	// Test cases
	tests := []struct {
		input  any
		output uint8
		err    error
	}{
		{10, 10, nil},
		{300, 0, go_easy_utils.ErrValOut},
		{"20", 20, nil},
		{1.5, 1, nil},
		{make(chan int), 0, go_easy_utils.ErrType},
	}

	// Test loop
	for _, test := range tests {
		result, err := AnyToUint8(test.input)
		if result != test.output || !reflect.DeepEqual(err, test.err) {
			t.Errorf("AnyToUint8(%v) = (%v, %v), expected (%v, %v)",
				test.input, result, err, test.output, test.err)
		}
	}
}

func TestAnyToUint16(t *testing.T) {
	// Test cases
	tests := []struct {
		input  any
		output uint16
		err    error
	}{
		{10, 10, nil},
		{70000, 0, go_easy_utils.ErrValOut},
		{"20", 20, nil},
		{1.5, 1, nil},
		{make(chan int), 0, go_easy_utils.ErrType},
	}

	// Test loop
	for _, test := range tests {
		result, err := AnyToUint16(test.input)
		if result != test.output || !reflect.DeepEqual(err, test.err) {
			t.Errorf("AnyToUint16(%v) = (%v, %v), expected (%v, %v)",
				test.input, result, err, test.output, test.err)
		}
	}
}

func TestAnyToUint32(t *testing.T) {
	// Test cases
	tests := []struct {
		input  any
		output uint32
		err    error
	}{
		{10, 10, nil},
		{5000000000, 0, go_easy_utils.ErrValOut},
		{"20", 20, nil},
		{1.5, 1, nil},
		{make(chan int), 0, go_easy_utils.ErrType},
	}

	// Test loop
	for _, test := range tests {
		result, err := AnyToUint32(test.input)
		if result != test.output || !reflect.DeepEqual(err, test.err) {
			t.Errorf("AnyToUint32(%v) = (%v, %v), expected (%v, %v)",
				test.input, result, err, test.output, test.err)
		}
	}
}

func TestAnyToUint64(t *testing.T) {
	iPtr := 90

	tests := []struct {
		name      string
		input     any
		want      uint64
		wantError error
	}{
		{
			name:      "Test -float32",
			input:     float32(-0.1),
			wantError: go_easy_utils.ErrUnsignedInt,
		},
		{
			name:      "Test -float64",
			input:     float64(-0.2),
			wantError: go_easy_utils.ErrUnsignedInt,
		},
		{
			name:      "Test -int",
			input:     int(-1),
			wantError: go_easy_utils.ErrUnsignedInt,
		},
		{
			name:      "Test -int8",
			input:     int8(-2),
			wantError: go_easy_utils.ErrUnsignedInt,
		},
		{
			name:      "Test -int16",
			input:     int16(-3),
			wantError: go_easy_utils.ErrUnsignedInt,
		},
		{
			name:      "Test -int32",
			input:     int32(-4),
			wantError: go_easy_utils.ErrUnsignedInt,
		},
		{
			name:      "Test -int64",
			input:     int64(-5),
			wantError: go_easy_utils.ErrUnsignedInt,
		},
		{
			name:  "Test uint",
			input: uint(12),
			want:  12,
		},
		{
			name:  "Test uint8",
			input: uint8(42),
			want:  42,
		},
		{
			name:  "Test uint16",
			input: uint16(42),
			want:  42,
		},
		{
			name:  "Test uint32",
			input: uint32(42),
			want:  42,
		},
		{
			name:  "Test uint64",
			input: uint64(42),
			want:  42,
		},
		{
			name:  "Test int8",
			input: int8(42),
			want:  42,
		},
		{
			name:  "Test int16",
			input: int16(42),
			want:  42,
		},
		{
			name:  "Test int32",
			input: int32(42),
			want:  42,
		},
		{
			name:  "Test int64",
			input: int64(42),
			want:  42,
		},
		{
			name:  "Test float32",
			input: float32(42.0),
			want:  42,
		},
		{
			name:  "Test float64",
			input: float64(42.0),
			want:  42,
		},
		{
			name:  "Test complex64",
			input: complex64(complex(42, 0)),
			want:  42,
		},
		{
			name:  "Test complex128",
			input: complex128(complex(42, 0)),
			want:  42,
		},
		{
			name:  "Test string",
			input: "42",
			want:  42,
		},
		{
			name:      "Test invalid string",
			input:     "not a number",
			wantError: go_easy_utils.ErrSyntax,
		},
		{
			name:      "Test nil pointer",
			input:     (*int)(nil),
			want:      0,
			wantError: nil,
		},
		{
			name:  "Test bool true",
			input: true,
			want:  1,
		},
		{
			name:  "Test bool false",
			input: false,
			want:  0,
		},
		{
			name:  "Test int point",
			input: &iPtr,
			want:  90,
		},
		{
			name:  "Test nil",
			input: nil,
			want:  0,
		},
		{
			name:      "test -complex",
			input:     complex(-1, -1),
			wantError: go_easy_utils.ErrUnsignedInt,
		},
		{
			name:      "Test invalid type",
			input:     make(chan int),
			wantError: go_easy_utils.ErrType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AnyToUint64(tt.input)
			if err != tt.wantError {
				t.Errorf("toUint64(%v) error = %v, wantError %v", tt.input, err, tt.wantError)
			}
			if got != tt.want {
				t.Errorf("toUint64(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
