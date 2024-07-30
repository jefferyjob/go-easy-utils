package cryptoUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "空字符串",
			input:    "",
			expected: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name:     "字符串",
			input:    "test",
			expected: "098f6bcd4621d373cade4e832627b4f6",
		},
		{
			name:     "带空格的字符串",
			input:    "hello world",
			expected: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Md5(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
