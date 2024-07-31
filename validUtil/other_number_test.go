package validUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPostalCode(t *testing.T) {
	testCases := []struct {
		name string
		code string
		want bool
	}{
		{"有效邮政编码1", "100000", true},
		{"有效邮政编码2", "200000", true},
		{"有效邮政编码3", "999999", true},
		{"无效邮政编码1", "1234567", false},
		{"无效邮政编码2", "2000000", false},
		{"无效邮政编码3", "123a56", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsPostalCode(tc.code)
			assert.Equal(t, tc.want, res)
		})
	}
}
