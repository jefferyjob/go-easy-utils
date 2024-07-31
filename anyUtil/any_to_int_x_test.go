package anyUtil

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAnyToInt(t *testing.T) {
	iPtr := 90
	testCases := []struct {
		name     string
		input    any
		expected int
		err      error
	}{
		// 测试整数输入
		{name: "number", input: 42, expected: 42, err: nil},
		{name: "int", input: int(42), expected: 42, err: nil},
		{name: "int64", input: int64(42), expected: 42, err: nil},
		{name: "int32", input: int32(42), expected: 42, err: nil},
		{name: "int", input: int(42), expected: 42, err: nil},
		{name: "int64", input: uint64(42), expected: 42, err: nil},
		{name: "int32", input: uint32(42), expected: 42, err: nil},
		{name: "uint", input: uint(42), expected: 42, err: nil},
		{name: "iPtr", input: &iPtr, expected: 90, err: nil},

		// 测试浮点数输入
		{name: "float64", input: float64(42.0), expected: 42, err: nil},
		{name: "float32", input: float32(42.0), expected: 42, err: nil},
		{name: "float64_小数点", input: float64(42.5), expected: 42, err: nil},

		// 测试非数值类型
		{name: "string", input: "rand string", expected: 0, err: ErrSyntax},
		{name: "nil", input: nil, expected: 0, err: nil},

		// 布尔
		{name: "true", input: true, expected: 1, err: nil},
		{name: "false", input: false, expected: 0, err: nil},

		// 空指针
		{name: "nil point", input: (*int)(nil), expected: 0, err: nil},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := AnyToInt(test.input)
			if res != test.expected || !errors.Is(err, test.err) {
				t.Errorf("AnyToInt(%v) = (%d, %v); want (%d, %v)", test.input, res, err, test.expected, test.err)
			}
		})
	}
}

func TestAnyToInt8(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected int8
		err      error
	}{
		// 测试整数输入
		{name: "number", input: 42, expected: 42, err: nil},
		{name: "int", input: int(42), expected: 42, err: nil},
		{name: "int64", input: int64(42), expected: 42, err: nil},
		{name: "int32", input: int32(42), expected: 42, err: nil},
		{name: "int", input: int(42), expected: 42, err: nil},
		{name: "int64", input: uint64(42), expected: 42, err: nil},
		{name: "int32", input: uint32(42), expected: 42, err: nil},
		{name: "uint", input: uint(42), expected: 42, err: nil},

		// 测试浮点数输入
		{name: "float64", input: float64(42.0), expected: 42, err: nil},
		{name: "float32", input: float32(42.0), expected: 42, err: nil},
		{name: "float64_小数点", input: float64(42.5), expected: 42, err: nil},

		// 测试非数值类型
		{name: "string", input: "rand string", expected: 0, err: ErrSyntax},
		{name: "nil", input: nil, expected: 0, err: nil},

		// 溢出
		{name: "MinInt8", input: math.MinInt8 - 1, expected: 0, err: ErrValOut},
		{name: "MaxInt8", input: math.MaxInt8 + 1, expected: 0, err: ErrValOut},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := AnyToInt8(test.input)
			if res != test.expected || !errors.Is(err, test.err) {
				t.Errorf("AnyToInt8(%v) = (%d, %v); want (%d, %v)", test.input, res, err, test.expected, test.err)
			}
		})
	}
}

func TestAnyToInt16(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected int16
		err      error
	}{
		// 测试整数输入
		{name: "number", input: 42, expected: 42, err: nil},
		{name: "int", input: int(42), expected: 42, err: nil},
		{name: "int64", input: int64(42), expected: 42, err: nil},
		{name: "int32", input: int32(42), expected: 42, err: nil},
		{name: "int", input: int(42), expected: 42, err: nil},
		{name: "int64", input: uint64(42), expected: 42, err: nil},
		{name: "int32", input: uint32(42), expected: 42, err: nil},
		{name: "uint", input: uint(42), expected: 42, err: nil},

		// 测试浮点数输入
		{name: "float64", input: float64(42.0), expected: 42, err: nil},
		{name: "float32", input: float32(42.0), expected: 42, err: nil},
		{name: "float64_小数点", input: float64(42.5), expected: 42, err: nil},

		// 测试非数值类型
		{name: "string", input: "rand string", expected: 0, err: ErrSyntax},
		{name: "nil", input: nil, expected: 0, err: nil},

		// 溢出
		{name: "MinInt16", input: math.MinInt16 - 1, expected: 0, err: ErrValOut},
		{name: "MaxInt16", input: math.MaxInt16 + 1, expected: 0, err: ErrValOut},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := AnyToInt16(test.input)
			if res != test.expected || !errors.Is(err, test.err) {
				t.Errorf("AnyToInt16(%v) = (%d, %v); want (%d, %v)", test.input, res, err, test.expected, test.err)
			}
		})
	}
}

func TestAnyToInt32(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected int32
		err      error
	}{
		// 测试整数输入
		{name: "number", input: 42, expected: 42, err: nil},
		{name: "int", input: int(42), expected: 42, err: nil},
		{name: "int64", input: int64(42), expected: 42, err: nil},
		{name: "int32", input: int32(42), expected: 42, err: nil},
		{name: "int", input: int(42), expected: 42, err: nil},
		{name: "int64", input: uint64(42), expected: 42, err: nil},
		{name: "int32", input: uint32(42), expected: 42, err: nil},
		{name: "uint", input: uint(42), expected: 42, err: nil},

		// 测试浮点数输入
		{name: "float64", input: float64(42.0), expected: 42, err: nil},
		{name: "float32", input: float32(42.0), expected: 42, err: nil},
		{name: "float64_小数点", input: float64(42.5), expected: 42, err: nil},

		// 测试非数值类型
		{name: "string", input: "rand string", expected: 0, err: ErrSyntax},
		{name: "nil", input: nil, expected: 0, err: nil},

		// 溢出
		{name: "MinInt32", input: math.MinInt32 - 1, expected: 0, err: ErrValOut},
		{name: "MaxInt32", input: math.MaxInt32 + 1, expected: 0, err: ErrValOut},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := AnyToInt32(test.input)
			if res != test.expected || !errors.Is(err, test.err) {
				t.Errorf("AnyToInt32(%v) = (%d, %v); want (%d, %v)", test.input, res, err, test.expected, test.err)
			}
		})
	}
}

func TestAnyToInt64(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected int64
		err      error
	}{
		// 测试整数输入
		{name: "number", input: 42, expected: 42, err: nil},
		{name: "int", input: int(42), expected: 42, err: nil},
		{name: "int64", input: int64(42), expected: 42, err: nil},
		{name: "int32", input: int32(42), expected: 42, err: nil},
		{name: "int", input: int(42), expected: 42, err: nil},
		{name: "int64", input: uint64(42), expected: 42, err: nil},
		{name: "int32", input: uint32(42), expected: 42, err: nil},
		{name: "uint", input: uint(42), expected: 42, err: nil},

		// 测试浮点数输入
		{name: "float64", input: float64(42.0), expected: 42, err: nil},
		{name: "float32", input: float32(42.0), expected: 42, err: nil},
		{name: "float64_小数点", input: float64(42.5), expected: 42, err: nil},

		// 测试非数值类型
		{name: "string", input: "rand string", expected: 0, err: ErrSyntax},
		{name: "nil", input: nil, expected: 0, err: nil},

		// 布尔
		{name: "true", input: true, expected: 1, err: nil},
		{name: "false", input: false, expected: 0, err: nil},

		// 空指针
		{name: "nil point", input: (*int)(nil), expected: 0, err: nil},

		// complex
		{name: "complex128", input: complex(3.14, 1.59), expected: 3, err: nil},
		{name: "complex64", input: complex(float32(2.71), float32(1.41)), expected: 2, err: nil},

		// 其他
		{name: "struct", input: struct{ Name string }{Name: "test"}, expected: 0, err: ErrType},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := AnyToInt64(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}
