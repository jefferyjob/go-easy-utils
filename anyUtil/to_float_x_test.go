package anyUtil

import (
	"testing"
)

func TestAnyToFloat32(t *testing.T) {
	cases := []struct {
		value    interface{}
		expected float32
		err      error
	}{
		{float32(1.23), float32(1.23), nil},
		{float64(2.34), float32(2.34), nil},
		{int(3), float32(3), nil},
		{int8(4), float32(4), nil},
		{int16(5), float32(5), nil},
		{int32(6), float32(6), nil},
		{int64(7), float32(7), nil},
		{uint(8), float32(8), nil},
		{uint8(9), float32(9), nil},
		{uint16(10), float32(10), nil},
		{uint32(11), float32(11), nil},
		{uint64(12), float32(12), nil},
		{"13.45", float32(13.45), nil},
	}

	for _, c := range cases {
		actual, err := AnyToFloat32(c.value)
		if actual != c.expected || err != c.err {
			t.Errorf("AnyToFloat32(%v) == (%v, %v), expected (%v, %v)", c.value, actual, err, c.expected, c.err)
		}
	}
}

func TestAnyToFloat64(t *testing.T) {
	cases := []struct {
		value    interface{}
		expected float64
		err      error
	}{
		{float64(2.34), float64(2.34), nil},
		{int(3), float64(3), nil},
		{int8(4), float64(4), nil},
		{int16(5), float64(5), nil},
		{int32(6), float64(6), nil},
		{int64(7), float64(7), nil},
		{uint(8), float64(8), nil},
		{uint8(9), float64(9), nil},
		{uint16(10), float64(10), nil},
		{uint32(11), float64(11), nil},
		{uint64(12), float64(12), nil},
		{"13.45", float64(13.45), nil},
	}

	for _, c := range cases {
		actual, err := AnyToFloat64(c.value)
		if actual != c.expected || err != c.err {
			t.Errorf("AnyToFloat64(%v) == (%v, %v), expected (%v, %v)", c.value, actual, err, c.expected, c.err)
		}
	}
}
