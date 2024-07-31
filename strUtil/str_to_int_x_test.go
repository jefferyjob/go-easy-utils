package strUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrToInt(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"零", "0", 0},
		{"正数", "1", 1},
		{"负数", "-1", -1},
		{"无效输入", "invalid", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToInt(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestStrToInt8(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int8
	}{
		{"零", "0", 0},
		{"正数", "1", 1},
		{"负数", "-1", -1},
		{"无效输入", "invalid", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToInt8(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestStrToInt16(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int16
	}{
		{"零", "0", 0},
		{"正数", "1", 1},
		{"负数", "-1", -1},
		{"无效输入", "invalid", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToInt16(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestStrToInt32(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int32
	}{
		{"零", "0", 0},
		{"正数", "1", 1},
		{"负数", "-1", -1},
		{"无效输入", "invalid", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToInt32(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestStrToInt64(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int64
	}{
		{"零", "0", 0},
		{"正数", "1", 1},
		{"负数", "-1", -1},
		{"最大正数", "9223372036854775807", 9223372036854775807},
		{"最小负数", "-9223372036854775808", -9223372036854775808},
		{"32位最大正数", "2147483647", 2147483647},
		{"32位最小负数", "-2147483648", -2147483648},
		{"无效输入", "invalid", 0},
		{"超出最大值", "9223372036854775808", 0},
		{"超出最小值", "-9223372036854775809", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToInt64(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
