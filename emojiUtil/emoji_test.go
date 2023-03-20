package emojiUtil

import (
	"strings"
	"testing"
)

func TestDecodeEmojiUnicode(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"[\\u1F602]", "😂"},
		{"[\\u1F60A]", "😊"},
		{"[\\u1F60E]456", "😎456"},
		{"No emoji", "No emoji"},
	}

	for _, tc := range testCases {
		res := DecodeEmojiUnicode(tc.input)

		// 转大写
		res = strings.ToUpper(res)
		expected := strings.ToUpper(tc.expected)

		if res != expected {
			t.Errorf("Unexpected result - input: %s, expected: %s, got: %s",
				tc.input,
				expected,
				res,
			)
		}
	}
}

func TestDecodeEmojiUnicode2(t *testing.T) {
	input := "[\\u1F602]"
	expected := "😂"
	result := DecodeEmojiUnicode(input)

	// 转大写
	expected = strings.ToUpper(expected)
	result = strings.ToUpper(result)

	if result != expected {
		t.Errorf("DecodeEmoji(%s) = %s; expected %s", input, result, expected)
	}
}

func TestEncodeEmojiUnicode(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Hello 😂", "Hello [\\u1F602]"},
		{"No emoji", "No emoji"},
	}

	for _, tc := range testCases {
		actual := EncodeEmojiUnicode(tc.input)

		// 转大写
		actual = strings.ToUpper(actual)
		expected := strings.ToUpper(tc.expected)

		if actual != expected {
			t.Errorf("Unexpected result - input: %s, expected: %s, got: %s",
				tc.input,
				expected,
				actual,
			)
		}
	}
}

func TestEncodeEmojiUnicode2(t *testing.T) {
	input := "😂"
	expected := "[\\u1F602]"
	result := EncodeEmojiUnicode(input)

	// 转大写
	expected = strings.ToUpper(expected)
	result = strings.ToUpper(result)

	if result != expected {
		t.Errorf("EncodeEmoji(%s) = %s; expected %s", input, result, expected)
	}
}
