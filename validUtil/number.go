package validUtil

import (
	"regexp"
)

// IsNumber 验证是否全部为数字
func IsNumber(input string) bool {
	reg := regexp.MustCompile("^[0-9]+$")
	return reg.MatchString(input)
}
