package anyUtil

import (
	"github.com/jefferyjob/go-easy-utils"
	"reflect"
	"testing"
)

func TestAnyToUint(t *testing.T) {
	// Test cases
	tests := []struct {
		input  interface{}
		output uint
		err    error
	}{
		{10, 10, nil},
		{-5, 0, go_easy_utils.ErrUnsignedInt},
		{"20", 20, nil},
		{1.5, 1, nil},
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
		input  interface{}
		output uint8
		err    error
	}{
		{10, 10, nil},
		{300, 0, go_easy_utils.ErrValOut},
		{"20", 20, nil},
		{1.5, 1, nil},
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
		input  interface{}
		output uint16
		err    error
	}{
		{10, 10, nil},
		{70000, 0, go_easy_utils.ErrValOut},
		{"20", 20, nil},
		{1.5, 1, nil},
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
		input  interface{}
		output uint32
		err    error
	}{
		{10, 10, nil},
		{5000000000, 0, go_easy_utils.ErrValOut},
		{"20", 20, nil},
		{1.5, 1, nil},
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
	// Test cases
	tests := []struct {
		input  interface{}
		output uint64
		err    error
	}{
		{10, 10, nil},
		{"20", 20, nil},
		{1.5, 1, nil},
	}

	// Test loop
	for _, test := range tests {
		result, err := AnyToUint64(test.input)
		if result != test.output || !reflect.DeepEqual(err, test.err) {
			t.Errorf("AnyToUint64(%v) = (%v, %v), expected (%v, %v)",
				test.input, result, err, test.output, test.err)
		}
	}
}
