package validUtil

import (
	"testing"
)

//func TestIsIDCard18Demo(t *testing.T) {
//	fmt.Println(IsIDCard18("120103199001015953"))
//	fmt.Println(IsIDCard18("44080319861221348X"))
//}

func TestIsIDCard(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"110102197809193026", true},
		{"142629680611101", true},
	}
	for _, c := range cases {
		got := IsIDCard(c.input)
		if got != c.want {
			t.Errorf("IsIDCard(%q) == %v, want %v", c.input, got, c.want)
		}
	}
}

func TestIsIDCard18(t *testing.T) {
	// 18位身份证号码测试
	cases18 := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"120103109001015953", false}, // 年份非法
		{"120103199018015953", false}, // 月份非法
		{"120103199001505953", false}, // 日期非法
		{"123456789012345678", false},
		{"12345678901234567X", false},
		{"12345678901234567x", false},
		{"110102197809193026", true},
		{"210381199503166219", true},
		{"64010219940313212X", true},
		{"120103199001015953", true},
		{"44080319861221348X", true},
	}
	for _, c := range cases18 {
		got := IsIDCard18(c.input)
		if got != c.want {
			t.Errorf("IsIDCard18(%q) == %v, want %v", c.input, got, c.want)
		}
	}
}

func TestIsIDCard15(t *testing.T) {
	// 15位身份证号码测试
	cases15 := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"142629680611101", true},
		{"610104620927690", true},
		{"142629601611101", false}, // 年份非法
		{"01345678901234", false},
		{"1234567890123X", false},
		{"9994567890123X", false},
	}
	for _, c := range cases15 {
		got := IsIDCard15(c.input)
		if got != c.want {
			t.Errorf("IsIDCard15(%q) == %v, want %v", c.input, got, c.want)
		}
	}
}
