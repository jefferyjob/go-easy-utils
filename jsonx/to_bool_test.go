package jsonx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
