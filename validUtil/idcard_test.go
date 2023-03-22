package validUtil

import "testing"

func TestIsIDCard1(t *testing.T) {
	// 18位身份证号码测试
	cases18 := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"123456789012345678", false},
		{"12345678901234567X", false},
		{"12345678901234567x", false},
		{"120103199001015953", true},
	}
	for _, c := range cases18 {
		got := IsIDCard(c.input)
		if got != c.want {
			t.Errorf("IsIDCard(%q) == %v, want %v", c.input, got, c.want)
		}
	}
}

func TestIsIDCard2(t *testing.T) {
	// 15位身份证号码测试
	cases15 := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"12345678901234", false},
		{"1234567890123X", false},
	}
	for _, c := range cases15 {
		got := IsIDCard(c.input)
		if got != c.want {
			t.Errorf("IsIDCard(%q) == %v, want %v", c.input, got, c.want)
		}
	}
}
