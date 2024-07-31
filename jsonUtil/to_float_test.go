package jsonUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
