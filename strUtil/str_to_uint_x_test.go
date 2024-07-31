package strUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrToUint(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected uint
	}{
		{"正数", "123", 123},
		{"零", "0", 0},
		{"负数", "-1", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToUint(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestStrToUint8(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected uint8
	}{
		{"正常值", "123", 123},
		{"零", "0", 0},
		{"负值", "-1", 0},
		{"超出范围", "256", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToUint8(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestStrToUint16(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected uint16
	}{
		{"正常值", "12345", 12345},
		{"零", "0", 0},
		{"负值", "-1", 0},
		{"超出范围", "65536", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToUint16(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestStrToUint32(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected uint32
	}{
		{"正常值", "1234567", 1234567},
		{"零", "0", 0},
		{"负值", "-1", 0},
		{"超出范围", "4294967296", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToUint32(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestStrToUint64(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput uint64
	}{
		{"正常值", "12345", 12345},
		{"零", "0", 0},
		{"最大值", "18446744073709551615", 18446744073709551615},
		{"负值", "-1", 0},
		{"非数字", "not a number", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToUint64(tc.input)
			assert.Equal(t, tc.expectedOutput, res)
		})
	}
}
