// Package emojiUtil emoji表情处理工具
package emojiUtil

import (
	"regexp"
	"strconv"
	"strings"
)

// DecodeEmojiUnicode Emoji表情解码
// 输入unicode编码，得到emoji表情
func DecodeEmojiUnicode(unicode string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(unicode, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			unicode = strings.Replace(unicode, src[i], string(rune(p)), -1)
		}
	}
	return unicode
}

// EncodeEmojiUnicode Emoji表情编码
// 输入emoji表情，得到unicode编码
func EncodeEmojiUnicode(emoji string) string {
	ret := ""
	rs := []rune(emoji)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u

		} else {
			ret += string(rs[i])
		}
	}
	return ret
}
