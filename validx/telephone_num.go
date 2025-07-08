package validx

import "regexp"

// IsMobile 验证是否为手机号码
func IsMobile(phone string) bool {
	match, _ := regexp.MatchString(`^1[3456789]\d{9}$`, phone)
	return match
}

// IsTelephone 验证是否为座机号码
func IsTelephone(telephone string) bool {
	match, _ := regexp.MatchString(`^0\d{2,3}-?\d{7,8}$`, telephone)
	return match
}
