package validUtil

import (
	"testing"
)

func TestIsNumber(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"123", true},
		{"1a2b3", false},
		{"0", true},
		{"", false},
		{" ", false},
	}
	for _, tc := range testCases {
		got := IsNumber(tc.input)
		if got != tc.want {
			t.Errorf("IsNumber(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
