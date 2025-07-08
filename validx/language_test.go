package validx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAllChinese(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"全中文", "中国人", true},
		{"全英文", "abc", false},
		{"中英文混合", "中abc", false},
		{"空字符串", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := IsAllChinese(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestIsContainChinese(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"全中文", "中国人", true},
		{"全英文", "abc", false},
		{"中英文混合", "中abc", true},
		{"空字符串", "", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := IsContainChinese(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestIsChineseName(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid chinese name", "张三", true},
		{"invalid chinese name", "abc", false},
		{"invalid chinese name", "张三张三张三张三", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsChineseName(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestIsEnglishName(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid english name", "John Smith", true},
		{"invalid english name", "John Smith 123", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsEnglishName(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
