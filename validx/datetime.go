package validx

import (
	"regexp"
	"strconv"
	"strings"
)

// IsTime 验证是否为时间格式（HH:mm:ss）
func IsTime(str string) bool {
	reg := regexp.MustCompile(`^(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d$`)
	return reg.MatchString(str)
}

// IsDate 验证是否为日期格式（YYYY-MM-DD）
func IsDate(str string) bool {
	reg := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !reg.MatchString(str) {
		return false
	}

	parts := strings.Split(str, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])
	if !isValidMonth(month) || !isValidDay(day, month, year) {
		return false
	}

	return true
}

// isValidMonth 验证月份是否合法 (1-12)
func isValidMonth(month int) bool {
	return month >= 1 && month <= 12
}

// isValidDay 验证天数是否合法
func isValidDay(day, month, year int) bool {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return day >= 1 && day <= 31
	case 4, 6, 9, 11:
		return day >= 1 && day <= 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
			return day >= 1 && day <= 29 // 闰年
		}
		return day >= 1 && day <= 28 // 平年
	default:
		return false
	}
}

// IsDateTime 验证是否为日期时间格式（yyyy-MM-dd HH:mm:ss）
func IsDateTime(str string) bool {
	reg := regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`)
	if !reg.MatchString(str) {
		return false
	}
	if !IsDate(str[0:10]) || !IsTime(str[11:]) {
		return false
	}
	return true
}
