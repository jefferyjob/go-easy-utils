package strUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrToBytes(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []byte
	}{
		{"转换hello", "hello", []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := StrToBytes(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
