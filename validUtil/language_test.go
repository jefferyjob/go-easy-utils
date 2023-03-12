package validUtil

import "testing"

func TestIsAllChinese(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"中国人", true},
		{"abc", false},
		{"中abc", false},
		{"", true},
	}

	for _, tt := range tests {
		got := IsAllChinese(tt.input)
		if got != tt.want {
			t.Errorf("IsAllChinese(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsContainChinese(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"中国人", true},
		{"abc", false},
		{"中abc", true},
		{"", false},
	}

	for _, tt := range tests {
		got := IsContainChinese(tt.input)
		if got != tt.want {
			t.Errorf("IsContainChinese(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsChineseName(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid chinese name", "张三", true},
		{"invalid chinese name", "abc", false},
		{"invalid chinese name", "张三张三张三张三", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := IsChineseName(tc.input); result != tc.expected {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestIsEnglishName(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid english name", "John Smith", true},
		{"invalid english name", "John Smith 123", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := IsEnglishName(tc.input); result != tc.expected {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
