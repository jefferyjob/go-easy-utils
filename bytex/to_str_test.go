package bytex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToStr(t *testing.T) {
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
			got := ToStr(tt.data)
			assert.Equal(t, tt.expected, got)
		})
	}
}
