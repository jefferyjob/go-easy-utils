// Package validUtil 验证银行卡号
package validUtil

import (
	"regexp"
)

// IsBankCardNo 验证是否为大陆银行卡号
func IsBankCardNo(cardNo string) bool {
	// 银行卡号长度为 16-19 位
	if len(cardNo) < 16 || len(cardNo) > 19 {
		return false
	}

	// 银行卡号必须为数字
	match, _ := regexp.MatchString("^[0-9]+$", cardNo)
	if !match {
		return false
	}

	// 使用 Luhn 算法验证银行卡号的合法性
	var sum int
	for i, c := range cardNo {
		num := int(c - '0')
		if i%2 == 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}
	return sum%10 == 0
}
