package validUtil

import (
	"testing"
)

// 测试IsBankCardNo函数
func TestIsBankCardNo(t *testing.T) {
	tests := []struct {
		name    string
		cardNo  string
		wantRes bool
	}{
		{"empty cardNo", "", false},
		{"less than 16 digits", "123456789012345", false},
		{"more than 19 digits", "12345678901234567890", false},
		{"contains non-numeric characters", "1234567a890123456", false},
		{"valid cardNo with 16 digits", "6222020123456789", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBankCardNo(tt.cardNo)
			if got != tt.wantRes {
				t.Errorf("IsBankCardNo(%s) = %t, want %t", tt.cardNo, got, tt.wantRes)
			}
		})
	}

	testCases := []struct {
		name     string
		cardNum  string
		expected bool
	}{
		{"case 1", "1234567890123456", false},
		{"case 2", "6217000010081698261", true},
		{"case 3", "6227000010081698261", false},
		{"case 4", "621700001008169826", false},
		{"case 5", "62170000100816982612", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsBankCardNo(tc.cardNum)
			if result != tc.expected {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
