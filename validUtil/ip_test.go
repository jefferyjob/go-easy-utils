package validUtil

import "testing"

func TestIsIPv4(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"192.168.0.1", true},
		{"10.0.0.1", true},
		{"172.16.0.1", true},
		{"1.2.3.4", true},
		{"1.2.3.4.5", false},
		{"1.2.3", false},
		{"abc.def.ghi.jkl", false},
	}

	for _, tt := range tests {
		got := IsIPv4(tt.input)
		if got != tt.want {
			t.Errorf("IsIPv4(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsIPv6(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"CDCD:910A:2222:5498:8475:1111:3900:2020", true},
		{"1030::C9B4:FF12:48AA:1A2B", true},
		{"2000:0:0:0:0:0:0:1", true},
		{"127.0.0.1", false},
		{"::ffff:127.0.0.1", false},
	}

	for _, tc := range testCases {
		result := IsIPv6(tc.input)
		if result != tc.expected {
			t.Errorf("Expected IsIPv6(%q) to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
