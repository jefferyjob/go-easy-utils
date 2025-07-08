package validx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDecimal(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"有效小数", "1.23", true},
		{"整数", "123", true},
		{"有效小数", "123.4", true},
		{"超长小数", "123.456", false},
		{"多个点", "1.2.3", false},
		{"非数字", "abc", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsDecimal(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}
