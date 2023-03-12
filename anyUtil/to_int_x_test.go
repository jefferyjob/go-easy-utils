package anyUtil

import (
	"errors"
	"fmt"
	"testing"
)

func TestAnyToInt(t *testing.T) {
	tests := []struct {
		input interface{}
		want  int
	}{
		{1, 1},
		{int8(2), 2},
		{int16(3), 3},
		{int32(4), 4},
		{int64(5), 5},
		{uint(6), 6},
		{uint8(7), 7},
		{uint16(8), 8},
		{uint32(9), 9},
		{uint64(10), 10},
		{"11", 11},
		{"invalid", 0},
	}
	for _, test := range tests {
		got, err := AnyToInt(test.input)
		if err != nil {
			if test.want != 0 {
				t.Errorf("AnyToInt(%v) error: %v", test.input, err)
			}
		} else if got != test.want {
			t.Errorf("AnyToInt(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestAnyToInt8(t *testing.T) {
	tests := []struct {
		input interface{}
		want  int8
		err   bool
	}{
		{1, 1, false},
		{int8(2), 2, false},
	}
	for _, test := range tests {
		got, err := AnyToInt8(test.input)
		if err != nil {
			if !test.err {
				t.Errorf("AnyToInt8(%v) error: %v", test.input, err)
			}
		} else if got != test.want {
			t.Errorf("AnyToInt8(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestAnyToInt16(t *testing.T) {
	tests := []struct {
		input interface{}
		want  int16
		err   error
	}{
		{int16(42), 42, nil},
		{int8(42), 42, nil},
		{int32(42), 42, nil},
		{int64(42), 42, nil},
		{uint(42), 42, nil},
		{uint8(42), 42, nil},
		{uint16(42), 42, nil},
		{uint32(42), 42, nil},
		{uint64(42), 42, nil},
		{"42", 42, nil},
	}

	for _, tt := range tests {
		got, err := AnyToInt16(tt.input)
		if got != tt.want || !errors.Is(err, tt.err) {
			t.Errorf("AnyToInt16(%v) = (%v, %v); want (%v, %v)", tt.input, got, err, tt.want, tt.err)
		}
	}
}

func TestAnyToInt32(t *testing.T) {
	tests := []struct {
		input interface{}
		want  int32
		err   error
	}{
		{int32(42), 42, nil},
		{int8(42), 42, nil},
		{int16(42), 42, nil},
	}

	for _, tt := range tests {
		got, err := AnyToInt32(tt.input)
		if got != tt.want || !errors.Is(err, tt.err) {
			t.Errorf("AnyToInt32(%v) = (%v, %v); want (%v, %v)", tt.input, got, err, tt.want, tt.err)
		}
	}
}

func TestAnyToInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int64
		err      error
	}{
		{"int", 42, 42, nil},
		{"int8", int8(42), 42, nil},
		{"int16", int16(42), 42, nil},
		{"int32", int32(42), 42, nil},
		{"int64", int64(42), 42, nil},
		{"uint", uint(42), 42, nil},
		{"uint8", uint8(42), 42, nil},
		{"uint16", uint16(42), 42, nil},
		{"uint32", uint32(42), 42, nil},
		{"uint64", uint64(42), 42, nil},
		{"string", "42", 42, nil},
		{"invalid", struct{}{}, 0, fmt.Errorf("unsupported type struct {}")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AnyToInt64(tt.input)
			if err != nil && err.Error() != tt.err.Error() {
				t.Fatalf("unexpected error: expected %v, got %v", tt.err, err)
			}
			if result != tt.expected {
				t.Fatalf("unexpected result: expected %v, got %v", tt.expected, result)
			}
		})
	}
}
