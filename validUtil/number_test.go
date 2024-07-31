package validUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNumber(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"纯数字", "123", true},
		{"含字母", "1a2b3", false},
		{"单个数字", "0", true},
		{"空字符串", "", false},
		{"空格", " ", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsNumber(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}
