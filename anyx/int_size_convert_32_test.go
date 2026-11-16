//go:build 386 || arm || mips || mipsle

package anyx

import (
	"errors"
	"math"
	"testing"
)

// TestToInt32BitPlatform 验证 32 位平台上 int 为 32 位时，
// ToInt 对超出范围的值返回 ErrValOut 而不是静默截断
func TestToInt32BitPlatform(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected int
		err      error
	}{
		{name: "MaxInt32+1 int64", input: int64(math.MaxInt32) + 1, expected: 0, err: ErrValOut},
		{name: "MinInt32-1 int64", input: int64(math.MinInt32) - 1, expected: 0, err: ErrValOut},
		{name: "MaxInt32+1 float64", input: float64(math.MaxInt32) + 1, expected: 0, err: ErrValOut},
		{name: "MinInt32-1 float64", input: float64(math.MinInt32) - 1, expected: 0, err: ErrValOut},
		{name: "超出范围字符串", input: "2147483648", expected: 0, err: ErrValOut},
		{name: "超出范围负数字符串", input: "-2147483649", expected: 0, err: ErrValOut},
		{name: "边界内 int64", input: int64(math.MaxInt32), expected: math.MaxInt32, err: nil},
		{name: "边界内字符串", input: "2147483647", expected: math.MaxInt32, err: nil},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := ToInt(test.input)
			if res != test.expected || !errors.Is(err, test.err) {
				t.Errorf("ToInt(%v) = (%d, %v); want (%d, %v)", test.input, res, err, test.expected, test.err)
			}
		})
	}
}

// TestToUint32BitPlatform 验证 32 位平台上 uint 为 32 位时，
// ToUint 对超出范围的值返回 ErrValOut 而不是静默截断
func TestToUint32BitPlatform(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected uint
		err      error
	}{
		{name: "MaxUint32+1 uint64", input: uint64(math.MaxUint32) + 1, expected: 0, err: ErrValOut},
		{name: "MaxUint32+1 float64", input: float64(math.MaxUint32) + 1, expected: 0, err: ErrValOut},
		{name: "超出范围字符串", input: "4294967296", expected: 0, err: ErrValOut},
		{name: "边界内 uint64", input: uint64(math.MaxUint32), expected: math.MaxUint32, err: nil},
		{name: "边界内字符串", input: "4294967295", expected: math.MaxUint32, err: nil},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := ToUint(test.input)
			if res != test.expected || !errors.Is(err, test.err) {
				t.Errorf("ToUint(%v) = (%d, %v); want (%d, %v)", test.input, res, err, test.expected, test.err)
			}
		})
	}
}
