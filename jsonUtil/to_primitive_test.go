package jsonUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	var iPar = "point"
	tests := []struct {
		name  string
		value any
		want  string
	}{
		{"nil", nil, ""},
		{"string", "hello", "hello"},
		{"int", 42, "42"},
		{"int8", int8(42), "42"},
		{"int16", int16(42), "42"},
		{"int32", int32(42), "42"},
		{"int64", int64(42), "42"},
		{"uint", uint(42), "42"},
		{"uint8", uint8(42), "42"},
		{"uint16", uint16(42), "42"},
		{"uint32", uint32(42), "42"},
		{"uint64", uint64(42), "42"},
		{"float32", float32(3.14159), "3.14159"},
		{"float64", 3.14159, "3.14159"},
		{"bool-true", true, "true"},
		{"bool-false", false, "false"},
		{"point", &iPar, "point"},
		{"complex64", complex64(1 + 2i), "(1+2i)"},
		{"complex128", complex128(3 + 4i), "(3+4i)"},
		{"chan", make(chan int), ""},
		{"int nil", (*int)(nil), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toStringReflect(tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAnyToInt64(t *testing.T) {
	var iPtr = 90
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
		{name: "point", input: &iPtr, expected: 90, err: nil},

		// complex
		{name: "complex128", input: complex(3.14, 1.59), expected: 3, err: nil},
		{name: "complex64", input: complex(float32(2.71), float32(1.41)), expected: 2, err: nil},

		// 其他
		{name: "struct", input: struct{ Name string }{Name: "test"}, expected: 0, err: ErrType},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := toInt64Reflect(tc.input)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestToUint64(t *testing.T) {
	var iPtr = 90
	tests := []struct {
		name      string
		input     any
		want      uint64
		wantError error
	}{
		{
			name:  "Test -float32",
			input: float32(-0.1),
			want:  0,
		},
		{
			name:  "Test -float64",
			input: float64(-0.2),
			want:  0,
		},
		{
			name:  "Test -int",
			input: int(-1),
			want:  0,
		},
		{
			name:  "Test -int8",
			input: int8(-2),
			want:  0,
		},
		{
			name:  "Test -int16",
			input: int16(-3),
			want:  0,
		},
		{
			name:  "Test -int32",
			input: int32(-4),
			want:  0,
		},
		{
			name:  "Test -int64",
			input: int64(-5),
			want:  0,
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
			name:  "test -complex",
			input: complex(-1, -1),
			want:  0,
		},
		{
			name:  "Test string",
			input: "42",
			want:  42,
		},
		{
			name:      "Test invalid string",
			input:     "not a number",
			wantError: ErrSyntax,
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
			name:      "Test invalid type",
			input:     make(chan int),
			wantError: ErrType,
		},
		{
			name:  "Test point",
			input: &iPtr,
			want:  90,
		},
		{
			name:  "nil",
			input: nil,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toUint64Reflect(tt.input)
			assert.Equal(t, tt.wantError, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToFloat64(t *testing.T) {
	var iPtr = 90
	testCases := []struct {
		name        string
		value       any
		expected    float64
		expectedErr error
	}{
		{"空值", nil, 0, nil},
		{"浮点数", float32(123.5), 123.5, nil},
		{"字符串数", "123.456", 123.456, nil},
		{"无符整型", uint(123), 123, nil},
		{"无符uint8", uint8(123), 123, nil},
		{"无符uint16", uint16(123), 123, nil},
		{"无符uint32", uint32(123), 123, nil},
		{"无符uint64", uint64(123), 123, nil},
		{"有符整型", int(123), 123, nil},
		{"有符int8", int8(123), 123, nil},
		{"有符int16", int16(123), 123, nil},
		{"有符int32", int32(123), 123, nil},
		{"有符int64", int64(123), 123, nil},
		{"复数64", complex64(1 + 2i), 1, nil},
		{"复数128", complex128(1 + 2i), 1, nil},
		{"布尔真", true, 1, nil},
		{"布尔假", false, 0, nil},
		{"空布尔指针", (*bool)(nil), 0, nil},
		{"非空指针", &iPtr, 90, nil},
		{"通道", make(chan int), 0, ErrType},
		{"无效字符串", "abc", 0, ErrSyntax},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := toFloat64Reflect(tc.value)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestToBool(t *testing.T) {
	var iPtr = 90
	var tests = []struct {
		name  string
		input any
		want  bool
	}{
		{"布尔真", true, true},
		{"布尔假", false, false},
		{"负整数", int(-1), true},
		{"正整数", int(1), true},
		{"零整数", int(0), false},
		{"正int8", int8(1), true},
		{"零int8", int8(0), false},
		{"正int16", int16(1), true},
		{"零int16", int16(0), false},
		{"正int32", int32(1), true},
		{"零int32", int32(0), false},
		{"正int64", int64(1), true},
		{"零int64", int64(0), false},
		{"正uint", uint(1), true},
		{"零uint", uint(0), false},
		{"正uint8", uint8(1), true},
		{"零uint8", uint8(0), false},
		{"正uint16", uint16(1), true},
		{"零uint16", uint16(0), false},
		{"正uint32", uint32(1), true},
		{"零uint32", uint32(0), false},
		{"正uint64", uint64(1), true},
		{"零uint64", uint64(0), false},
		{"浮点1.0", float32(1.0), true},
		{"浮点0.0", float32(0.0), false},
		{"双精1.0", float64(1.0), true},
		{"双精0.0", float64(0.0), false},
		{"字符串", "abc", true},
		{"字符串真", "true", true},
		{"字符串假", "false", false},
		{"空字符串", "", false},
		{"空值", nil, false},
		{"非空指针", &iPtr, true},
		{"复数1+1i", complex64(1 + 1i), true},
		{"复数0+0i", complex64(0 + 0i), false},
		{"双复1+1i", complex128(1 + 1i), true},
		{"双复0+0i", complex128(0 + 0i), false},
		{"空指针", (*int)(nil), false},
		{"通道", make(chan int), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := toBoolReflect(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}
