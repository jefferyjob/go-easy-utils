package validx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMobile(t *testing.T) {
	type test struct {
		name  string
		input string
		want  bool
	}

	tests := []test{
		{"有效号码1", "15812345678", true},
		{"有效号码2", "15912345678", true},
		{"有效号码3", "17012345678", true},
		{"有效号码4", "17112345678", true},
		{"有效号码5", "17212345678", true},
		{"有效号码6", "18912345678", true},
		{"无效号码1", "29012345678", false},
		{"无效号码2", "11111111111", false},
		{"无效号码3", "09212345678", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := IsMobile(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestIsTelephone(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"有效号码1", "010-12345678", true},
		{"有效号码2", "02112345678", true},
		{"有效号码3", "075512345678", true},
		{"无效号码1", "12345678", false},
		{"无效号码2", "010-1234-5678", false},
		{"无效号码3", "0515123456a", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsTelephone(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}
