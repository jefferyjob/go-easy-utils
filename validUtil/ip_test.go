package validUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsIPv4(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"有效地址1", "192.168.0.1", true},
		{"有效地址2", "10.0.0.1", true},
		{"有效地址3", "172.16.0.1", true},
		{"有效地址4", "1.2.3.4", true},
		{"无效地址1", "1.2.3.4.5", false},
		{"无效地址2", "1.2.3", false},
		{"无效地址3", "abc.def.ghi.jkl", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := IsIPv4(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestIsIPv6(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"有效地址1", "CDCD:910A:2222:5498:8475:1111:3900:2020", true},
		{"有效地址2", "1030::C9B4:FF12:48AA:1A2B", true},
		{"有效地址3", "2000:0:0:0:0:0:0:1", true},
		{"无效地址1", "127.0.0.1", false},
		{"无效地址2", "::ffff:127.0.0.1", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsIPv6(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
