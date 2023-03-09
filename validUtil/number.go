package validUtil

import "regexp"

// IsMobile 验证是否为手机号码
func IsMobile(mobileNum string) bool {
	var regular = "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// IsTelephone 验证是否为座机号码
func IsTelephone(telephone string) bool {
	match, _ := regexp.MatchString(`^0\d{2,3}-?\d{7,8}$`, telephone)
	return match
}

// IsPostalCode 验证是否为邮编号码
func IsPostalCode(str string) bool {
	reg := regexp.MustCompile(`^[1-9]\d{5}$`)
	return reg.MatchString(str)
}

// IsDecimal 验证给定的字符串小数点后是否最多两位
func IsDecimal(input string) bool {
	// 定义小数的正则表达式，小数点后最多两位
	pattern := `^[0-9]+(\.[0-9]{1,2})?$`
	match, err := regexp.MatchString(pattern, input)
	return match && err == nil
}

// IsNumber 验证是否全部为数字
func IsNumber(input string) bool {
	reg := regexp.MustCompile("^[0-9]+$")
	return reg.MatchString(input)
}
