package emojiUtil

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDecodeEmojiUnicode(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"表情笑哭", "[\\u1F602]", "😂"},
		{"表情微笑", "[\\u1F60A]", "😊"},
		{"表情和数字", "[\\u1F60E]456", "😎456"},
		{"无意义的字符串", "No emoji", "No emoji"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := DecodeEmojiUnicode(tc.input)
			// 转大写
			res = strings.ToUpper(res)
			expected := strings.ToUpper(tc.expected)
			assert.Equal(t, expected, res)
		})
	}
}

func TestEncodeEmojiUnicode(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"字符串和表情", "Hello 😂", "Hello [\\u1F602]"},
		{"无意义的字符串", "No emoji", "No emoji"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := EncodeEmojiUnicode(tc.input)
			// 转大写
			actual = strings.ToUpper(actual)
			expected := strings.ToUpper(tc.expected)
			assert.Equal(t, expected, actual)
		})
	}
}
