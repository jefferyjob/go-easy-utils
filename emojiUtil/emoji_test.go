package emojiUtil

//func TestUnicodeEmojiDecode(t *testing.T) {
//	testCases := []struct {
//		input    string
//		expected string
//	}{
//		{"Hello [\\ud83d\\ude0a] [\\ud83d\\ude0b]", "Hello 😊 😋"},
//		{"[\\ud83d\\ude0a]", "😊"},
//		{"[\\ud83d\\ude0a]123[\\ud83d\\ude0b]456", "😊123😋456"},
//		{"No emoji", "No emoji"},
//	}
//
//	for _, tc := range testCases {
//		actual := UnicodeEmojiDecode(tc.input)
//		if actual != tc.expected {
//			t.Errorf("Unexpected result - input: %s, expected: %s, got: %s", tc.input, tc.expected, actual)
//		}
//	}
//}

//func TestUnicodeEmojiCode(t *testing.T) {
//	testCases := []struct {
//		input    string
//		expected string
//	}{
//		{"Hello 😊 😋", "Hello [\\ud83d\\ude0a] [\\ud83d\\ude0b]"},
//		{"😊", "[\\ud83d\\ude0a]"},
//		{"😊123😋456", "[\\ud83d\\ude0a]123[\\ud83d\\ude0b]456"},
//		{"No emoji", "No emoji"},
//	}
//
//	for _, tc := range testCases {
//		actual := UnicodeEmojiCode(tc.input)
//		if actual != tc.expected {
//			t.Errorf("Unexpected result - input: %s, expected: %s, got: %s", tc.input, tc.expected, actual)
//		}
//	}
//}
