package validUtil

//import "testing"
//
//func TestIsIDCard(t *testing.T) {
//	// 18位身份证号码测试
//	cases18 := []struct {
//		input string
//		want  bool
//	}{
//		{"", false},
//		{"123456789012345678", false},
//		{"12345678901234567X", false},
//		{"12345678901234567x", false},
//		{"360313198909085215", true},
//		{"36031319890908521x", true},
//		{"36031319890908521X", true},
//		{"110101198001010018", true},
//		{"11010119800101001X", true},
//		{"11010119800101001x", true},
//		{"320311770706001", false},
//		{"32031119770706001X", false},
//		{"360313198909285215", false},
//		{"360313198909085211", false},
//		{"360313198912345678", false},
//	}
//	for _, c := range cases18 {
//		got := IsIDCard(c.input)
//		if got != c.want {
//			t.Errorf("IsIDCard(%q) == %v, want %v", c.input, got, c.want)
//		}
//	}
//
//	// 15位身份证号码测试
//	cases15 := []struct {
//		input string
//		want  bool
//	}{
//		{"", false},
//		{"12345678901234", false},
//		{"1234567890123X", false},
//		{"1234567890123x", false},
//		{"320311770706001", true},
//		{"330782820506908", true},
//		{"620321871001001", true},
//		{"622922780623068", true},
//		{"632523830310017", true},
//		{"130283820529063", true},
//		{"340702800818200", true},
//		{"110000400101001", true},
//		{"110000500101001", false},
//		{"110000200101001", false},
//		{"110000700101001", false},
//	}
//	for _, c := range cases15 {
//		got := IsIDCard(c.input)
//		if got != c.want {
//			t.Errorf("IsIDCard(%q) == %v, want %v", c.input, got, c.want)
//		}
//	}
//}
