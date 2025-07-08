package validx

import (
	"regexp"
	"unicode"
)

// IsAllChinese 验证给定的字符串全部为中文
func IsAllChinese(input string) bool {
	for _, r := range input {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			return false
		}
	}
	return true
}

// IsContainChinese 验证给定的字符串包含中文
func IsContainChinese(input string) bool {
	for _, r := range input {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// IsChineseName 验证是否为中文名
func IsChineseName(name string) bool {
	pattern := "^[\u4E00-\u9FA5]{2,6}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(name)
}

// IsEnglishName 验证是否为英文名
func IsEnglishName(name string) bool {
	match, _ := regexp.MatchString(`^([a-zA-Z]+\s)*[a-zA-Z]+$`, name)
	return match
}
