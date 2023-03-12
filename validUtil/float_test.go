package validUtil

import "testing"

func TestIsDecimal(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"1.23", true},
		{"123", true},
		{"123.4", true},
		{"123.456", false},
		{"1.2.3", false},
		{"abc", false},
	}
	for _, tc := range testCases {
		got := IsDecimal(tc.input)
		if got != tc.want {
			t.Errorf("IsDecimal(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
