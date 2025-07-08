package validx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsIDCard(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"空字符串", "", false},
		{"有效身份证18位", "110102197809193026", true},
		{"有效身份证15位", "142629680611101", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsIDCard(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestIsIDCard18(t *testing.T) {
	cases18 := []struct {
		name  string
		input string
		want  bool
	}{
		{"空字符串", "", false},
		{"年份小于1900错误", "120103180001015953", false},
		{"年份非法", "120103109001015953", false},
		{"月份非法", "120103199018015953", false},
		{"月份非法", "622826199018015953", false},
		{"日期非法", "120103199001505953", false},
		{"无效号码1", "123456789012345678", false},
		{"无效号码2", "12345678901234567X", false},
		{"无效号码3", "12345678901234567x", false},
		{"有效号码1", "110102197809193026", true},
		{"有效号码2", "210381199503166219", true},
		{"有效号码3", "64010219940313212X", true},
		{"有效号码4", "120103199001015953", true},
		{"有效号码5", "44080319861221348X", true},
	}

	for _, tc := range cases18 {
		t.Run(tc.name, func(t *testing.T) {
			res := IsIDCard18(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestIsIDCard15(t *testing.T) {
	cases15 := []struct {
		name  string
		input string
		want  bool
	}{
		{"空字符串", "", false},
		{"有效号码1", "142629680611101", true},
		{"有效号码2", "610104620927690", true},
		{"年份非法", "142629601611101", false},
		{"无效号码1", "01345678901234", false},
		{"无效号码2", "1234567890123X", false},
		{"无效号码3", "9994567890123X", false},
		{"无效号码4", "102629680611101", false},
	}

	for _, tc := range cases15 {
		t.Run(tc.name, func(t *testing.T) {
			res := IsIDCard15(tc.input)
			assert.Equal(t, tc.want, res)
		})
	}
}
