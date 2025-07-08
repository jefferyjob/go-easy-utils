package emojix

import (
	"regexp"
	"strconv"
	"strings"
)

// Decode Emoji表情解码
// 输入unicode编码，得到emoji表情
func Decode(unicode string) string {
	// emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	// 提取emoji数据表达式
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

// Encode Emoji表情编码
// 输入emoji表情，得到unicode编码
func Encode(emoji string) string {
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
