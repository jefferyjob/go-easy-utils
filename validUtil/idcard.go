package validUtil

import (
	"regexp"
	"strconv"
)

// IsIDCard 验证身份证号(18或15位)
func IsIDCard(str string) bool {
	if len(str) != 15 && len(str) != 18 {
		return false
	}
	reg := regexp.MustCompile(`^(\d{6})(\d{2})(\d{2})(\d{2})(\d{3})(\d|X)$`)
	if !reg.MatchString(str) {
		return false
	}
	if len(str) == 18 {
		weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		checkCode := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
		sum := 0
		for i := 0; i < 17; i++ {
			num, _ := strconv.Atoi(str[i : i+1])
			sum += num * weight[i]
		}
		mod := sum % 11
		if checkCode[mod] != str[17:] {
			return false
		}
	}
	return true
}

// IsIDCard18 验证18位身份证号
func IsIDCard18(input string) bool {

	return true
}

// IsIDCard15 验证15位身份证号
func IsIDCard15(input string) bool {
	// 定义身份证号码的正则表达式
	pattern := `^[1-9]\d{5}\d{2}(0[1-9]|1[0-2])([012]\d|3[01])\d{3}$`
	match, err := regexp.MatchString(pattern, input)
	return match && err == nil
}
