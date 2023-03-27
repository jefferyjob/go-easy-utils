package anyUtil

import (
	go_easy_utils "github.com/jefferyjob/go-easy-utils"
	"math"
	"testing"
)

//func TestAnyToFloat64Test(t *testing.T) {
//	fmt.Println(AnyToFloat64([]int{}))
//}

func TestAnyToFloat32(t *testing.T) {
	tests := []struct {
		input interface{}
		want  float32
		err   error
	}{
		{nil, 0, nil},
		{float32(3.14), 3.14, nil},
		{float64(6.28), 6.28, nil},
		{int(42), 42, nil},
		{int8(8), 8, nil},
		{int16(16), 16, nil},
		{int32(32), 32, nil},
		{int64(64), 64, nil},
		{uint(24), 24, nil},
		{uint8(8), 8, nil},
		{uint16(16), 16, nil},
		{uint32(32), 32, nil},
		{uint64(64), 64, nil},
		{"3.14", 3.14, nil},
		{math.MaxFloat64, 0, go_easy_utils.ErrValOut},
		{make(chan int), 0, go_easy_utils.ErrType},
	}

	for _, test := range tests {
		got, err := AnyToFloat32(test.input)
		if got != test.want {
			t.Errorf("AnyToFloat32(%v) = %v; want %v", test.input, got, test.want)
		}
		if err != test.err {
			if err != nil && test.err != nil {
				if err.Error() != test.err.Error() {
					t.Errorf("AnyToFloat32(%v) error = %v; want %v", test.input, err, test.err)
				}
			} else {
				t.Errorf("AnyToFloat32(%v) error = %v; want %v", test.input, err, test.err)
			}
		}
	}
}

func TestAnyToFloat64(t *testing.T) {
	iPtr := 90

	tests := []struct {
		input interface{}
		want  float64
		err   error
	}{
		{nil, 0, nil},
		//{float32(3.14), 3.14, nil},
		{float64(6.28), 6.28, nil},
		{int(42), 42, nil},
		{int8(8), 8, nil},
		{int16(16), 16, nil},
		{int32(32), 32, nil},
		{int64(64), 64, nil},
		{uint(24), 24, nil},
		{uint8(8), 8, nil},
		{uint16(16), 16, nil},
		{uint32(32), 32, nil},
		{uint64(64), 64, nil},
		{"3.14", 3.14, nil},
		{true, 1, nil},
		{false, 0, nil},
		{complex64(1 + 2i), 1, nil},
		{complex128(1 + 2i), 1, nil},
		{&iPtr, 90, nil},
		{(*int)(nil), 0, nil},
		{"abc", 0, go_easy_utils.ErrSyntax},
		{[]int{}, 0, go_easy_utils.ErrType},
	}

	for _, test := range tests {
		got, err := AnyToFloat64(test.input)
		if got != test.want {
			t.Errorf("AnyToFloat64(%v) = %v; want %v", test.input, got, test.want)
		}
		if err != test.err {
			if err != nil && test.err != nil {
				if err.Error() != test.err.Error() {
					t.Errorf("AnyToFloat64(%v) error = %v; want %v", test.input, err, test.err)
				}
			} else {
				t.Errorf("AnyToFloat64(%v) error = %v; want %v", test.input, err, test.err)
			}
		}
	}
}
