package validx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBankCardNo(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty cardNo", "", false},
		{"less than 16 digits", "123456789012345", false},
		{"more than 19 digits", "12345678901234567890", false},
		{"contains non-numeric characters", "1234567a890123456", false},
		{"valid cardNo with 16 digits", "6222020123456789", true},

		{"case 1", "1234567890123456", false},
		{"case 2", "6217000010081698261", true},
		{"case 3", "6227000010081698261", false},
		{"case 4", "621700001008169826", false},
		{"case 5", "62170000100816982612", false},

		{"存在字母的银行卡号", "A217000010081698261", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsBankCardNo(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}
