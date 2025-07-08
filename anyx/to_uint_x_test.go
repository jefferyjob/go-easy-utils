package anyx

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestToUint(t *testing.T) {
	iPtr := 90
	testCases := []struct {
		name     string
		input    any
		expected uint
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

		// complex
		{name: "complex128", input: complex(3.14, 1.59), expected: 3, err: nil},
		{name: "complex64", input: complex(float32(2.71), float32(1.41)), expected: 2, err: nil},

		// 其他
		{name: "struct", input: struct{ Name string }{Name: "test"}, expected: 0, err: ErrType},

		{
			name:     "Value within uint range (32-bit)",
			input:    uint32(4294967295), // Maximum value for uint32 (32-bit)
			expected: uint(4294967295),
			err:      nil,
		},
		{
			name:     "Value within uint range (64-bit)",
			input:    uint64(18446744073709551615), // Maximum value for uint64 (64-bit)
			expected: uint(18446744073709551615),   // This is an edge case for a 64-bit system
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ToUint(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestToUint8(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected uint8
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
		{name: "MaxUint8", input: math.MaxUint8 + 1, expected: 0, err: ErrValOut},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ToUint8(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestToUint16(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected uint16
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
		{name: "MaxUint16", input: math.MaxUint16 + 1, expected: 0, err: ErrValOut},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ToUint16(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestToUint32(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected uint32
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
		{name: "MaxUint32", input: math.MaxUint32 + 1, expected: 0, err: ErrValOut},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ToUint32(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestToUint64(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected uint64
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

		// complex
		{name: "complex128", input: complex(3.14, 1.59), expected: 3, err: nil},
		{name: "complex64", input: complex(float32(2.71), float32(1.41)), expected: 2, err: nil},

		// 其他
		{name: "struct", input: struct{ Name string }{Name: "test"}, expected: 0, err: ErrType},

		// 小于0
		{name: "int < 0", input: int(-8), expected: 0, err: ErrUnsignedInt},
		{name: "float < 0", input: float64(-8), expected: 0, err: ErrUnsignedInt},
		{name: "complex128", input: complex(-3.14, -1.59), expected: 0, err: ErrUnsignedInt},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ToUint64(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}
