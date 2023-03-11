package validUtil

import (
	"encoding/json"
	"net"
	"net/url"
	"regexp"
	"strconv"
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

// IsEmail 是否为email
func IsEmail(input string) bool {
	// 定义邮箱地址的正则表达式
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, input)
	return match && err == nil
}

// IsIPv4 是否为ipv4地址
func IsIPv4(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil && ip.To4() != nil
}

// IsIPv6 是否为ipv6地址
func IsIPv6(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil && ip.To4() == nil
}

// IsURL 是否为URL地址
func IsURL(input string) bool {
	_, err := url.Parse(input)
	return err == nil
}

// IsJSON 是否为Json
func IsJSON(input string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(input), &js) == nil
}

// IsBankCardNo 验证是否为银行卡号
func IsBankCardNo(str string) bool {
	if len(str) < 15 || len(str) > 19 {
		return false
	}
	reg := regexp.MustCompile(`^\d{15,19}$`)
	if !reg.MatchString(str) {
		return false
	}
	sum := 0
	for i := len(str) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(str[i : i+1])
		if (len(str)-i)%2 == 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}
	return sum%10 == 0
}

// IsChineseName 验证是否为中文名
func IsChineseName(name string) bool {
	match, _ := regexp.MatchString(`^([\u4e00-\u9fa5]){2,7}$`, name)
	return match
}

// IsEnglishName 验证是否为英文名
func IsEnglishName(name string) bool {
	match, _ := regexp.MatchString(`^([a-zA-Z]+\s)*[a-zA-Z]+$`, name)
	return match
}

// IsQQ 验证是否为QQ号
func IsQQ(qq string) bool {
	match, _ := regexp.MatchString(`^[1-9][0-9]{4,10}$`, qq)
	return match
}

// IsWeChat 验证是否为微信号
func IsWeChat(wechat string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z][-_a-zA-Z0-9]{5,19}$`, wechat)
	return match
}

// IsWeibo 验证是否为微博ID
func IsWeibo(weibo string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9_-]{1,50}$`, weibo)
	return match
}

// IsPassword 验证密码是否合法
// 密码长度在6-20个字符之间，只包含数字、字母和下划线
func IsPassword(password string) bool {
	pattern := `^\w{6,20}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(password)
}
