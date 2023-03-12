// Package validUtil 验证社交账号
package validUtil

import (
	"encoding/json"
	"regexp"
)

// IsEmail 是否为email
func IsEmail(input string) bool {
	// 定义邮箱地址的正则表达式
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, input)
	return match && err == nil
}

// IsJSON 是否为Json
func IsJSON(input string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(input), &js) == nil
}

// IsQQ 验证是否为QQ号
func IsQQ(qq string) bool {
	match, _ := regexp.MatchString(`^[1-9][0-9]{4,12}$`, qq)
	return match
}

// IsWeChat 验证是否为微信号
func IsWeChat(wechat string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z][-_a-zA-Z0-9]{6,20}$`, wechat)
	return match
}

// IsWeibo 验证是否为微博ID
func IsWeibo(weibo string) bool {
	if len(weibo) < 6 || len(weibo) > 20 {
		return false
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z][\w-]*$`, weibo); !matched {
		return false
	}

	return true
}

// IsPassword 验证密码是否合法
// 密码长度在6-20个字符之间，必须包含数字、字母和特殊符号
func IsPassword(password string) bool {
	if len(password) < 6 || len(password) > 20 {
		return false
	}

	if matched, _ := regexp.MatchString(`[a-zA-Z]`, password); !matched {
		return false
	}

	if matched, _ := regexp.MatchString(`\d`, password); !matched {
		return false
	}

	if matched, _ := regexp.MatchString(`[^a-zA-Z\d]`, password); !matched {
		return false
	}

	return true
}
