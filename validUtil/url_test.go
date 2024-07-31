package validUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsURL(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"有效网址1", "https://www.google.com", true},
		{"有效网址2", "http://example.com", true},
		{"有效网址3", "https://www.example.com/path?query=string", true},
		{"无效网址1", "ftp://ftp.example.com", false},
		{"无效网址2", "invalid url", false},
		{"无效网址3", "www.example.com", false},
		{"无效网址4", "example.com", false},
		{"无效网址5", "http://", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsURL(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
