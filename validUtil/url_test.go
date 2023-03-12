package validUtil

import (
	"testing"
)

func TestIsURL(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"https://www.google.com", true},
		{"http://example.com", true},
		{"https://www.example.com/path?query=string", true},
		{"ftp://ftp.example.com", false},
		{"invalid url", false},
		{"www.example.com", false},
		{"example.com", false},
		{"http://", false},
	}

	for _, tc := range testCases {
		result := IsURL(tc.input)
		if result != tc.expected {
			t.Errorf("Expected IsURL(%q) to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
