package validUtil

import "regexp"

// IsPostalCode 验证是否为邮编号码
func IsPostalCode(str string) bool {
	reg := regexp.MustCompile(`^[1-9]\d{5}$`)
	return reg.MatchString(str)
}
