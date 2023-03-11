package validUtil

import (
	"regexp"
	"strconv"
)

// IsTime 验证是否为时间格式（HH:mm:ss）
func IsTime(str string) bool {
	reg := regexp.MustCompile(`^(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d$`)
	return reg.MatchString(str)
}

// IsDate 验证是否为日期格式（yyyy-MM-dd）
func IsDate(str string) bool {
	reg := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !reg.MatchString(str) {
		return false
	}
	year, _ := strconv.Atoi(str[0:4])
	month, _ := strconv.Atoi(str[5:7])
	day, _ := strconv.Atoi(str[8:10])
	if month < 1 || month > 12 {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}
	if month == 2 {
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			if day > 29 {
				return false
			}
		} else {
			if day > 28 {
				return false
			}
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		if day > 30 {
			return false
		}
	}
	return true
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
