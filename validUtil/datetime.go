package validUtil

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

// IsDate 验证是否为日期格式（yyyy-MM-dd）
func IsDate(str string) bool {
	reg := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !reg.MatchString(str) {
		return false
	}
	parts := strings.Split(str, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])
	return isValidMonth(month, year) && isValidDay(day, month, year)
}

func isValidMonth(month, year int) bool {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return true
	case 4, 6, 9, 11:
		return false
	case 2:
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			return true
		}
		return false
	}
	return false
}

func isValidDay(day, month, year int) bool {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day >= 1 && day <= 31 {
			return true
		}
	case 4, 6, 9, 11:
		if day >= 1 && day <= 30 {
			return true
		}
	case 2:
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			if day >= 1 && day <= 29 {
				return true
			}
		} else {
			if day >= 1 && day <= 28 {
				return true
			}
		}
	}
	return false
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
