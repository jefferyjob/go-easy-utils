package jsonUtil

import (
	"github.com/jefferyjob/go-easy-utils"
	"testing"
)

func TestToInt64(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected int64
		err      error
	}{
		{"123", 123, nil},
		{"-234", -234, nil},
		{345, 345, nil},
		{-456, -456, nil},
		{int8(12), 12, nil},
		{int16(1238), 1238, nil},
		{int32(1239), 1239, nil},
		{int64(1230), 1230, nil},
		{uint(1231), 1231, nil},
		{uint8(123), 123, nil},
		{uint16(1232), 1232, nil},
		{uint32(1233), 1233, nil},
		{uint64(1234), 1234, nil},
		{float32(12.45), 12, nil},
		{float64(123.45), 123, nil},
		{true, 1, nil},
		{false, 0, nil},
		{complex64(1 + 2i), 1, nil},
		{complex128(1 + 2i), 1, nil},
		{nil, 0, nil},
		{"not a number", 0, go_easy_utils.ErrSyntax},
		{make(chan int), 0, go_easy_utils.ErrType},
	}

	for _, tt := range tests {
		actual, err := toInt64(tt.input)
		if err != tt.err {
			t.Errorf("toInt64(%v): expected error %v, actual error %v", tt.input, tt.err, err)
		}
		if actual != tt.expected {
			t.Errorf("toInt64(%v): expected %v, actual %v", tt.input, tt.expected, actual)
		}

		actual, err = toInt64Reflect(tt.input)
		if err != tt.err {
			t.Errorf("toInt64(%v): expected error %v, actual error %v", tt.input, tt.err, err)
		}
		if actual != tt.expected {
			t.Errorf("toInt64(%v): expected %v, actual %v", tt.input, tt.expected, actual)
		}
	}
}
