package byteUtil

import "testing"

func TestBytesToStr(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected string
	}{
		{"empty slice", []byte{}, ""},
		{"non-empty slice", []byte{'h', 'e', 'l', 'l', 'o'}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToStr(tt.data)
			if got != tt.expected {
				t.Errorf("BytesToStr(%v) = %v, want %v", tt.data, got, tt.expected)
			}
		})
	}
}
