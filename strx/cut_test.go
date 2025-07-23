package strx

import (
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		sub      string
		n        int
		expected string
	}{
		{"Remove all", "hello world world", "world", 0, "hello  "},
		{"Remove once", "hello world world", "world", 1, "hello  world"},
		{"No sub", "hello world", "", 0, "hello world"},
		{"No match", "hello world", "abc", 0, "hello world"},
		{"Empty string", "", "world", 0, ""},
		{"Remove twice max", "go go go", "go", 2, "  go"},
		{"Remove all with zero count", "go go go", "go", 0, "  "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result string
			if tt.n > 0 {
				result = Cut(tt.s, tt.sub, tt.n)
			} else {
				result = Cut(tt.s, tt.sub)
			}
			if result != tt.expected {
				t.Errorf("Cut(%q, %q, %v) = %q; want %q", tt.s, tt.sub, tt.n, result, tt.expected)
			}
		})
	}
}
