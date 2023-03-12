package validUtil

import "testing"

func TestIsMobile(t *testing.T) {
	type test struct {
		input string
		want  bool
	}
	tests := []test{
		{input: "15812345678", want: true},
		{input: "15912345678", want: true},
		{input: "17012345678", want: true},
		{input: "17112345678", want: true},
		{input: "17212345678", want: true},
		{input: "18912345678", want: true},
		{input: "29012345678", want: false},
		{input: "11111111111", want: false},
		{input: "09212345678", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := IsMobile(tt.input); got != tt.want {
				t.Errorf("IsMobile(%s) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsTelephone(t *testing.T) {
	// Test valid telephone numbers
	validNumbers := []string{"010-12345678", "02112345678", "075512345678"}
	for _, num := range validNumbers {
		if !IsTelephone(num) {
			t.Errorf("%s should be a valid telephone number", num)
		}
	}

	// Test invalid telephone numbers
	invalidNumbers := []string{"12345678", "010-1234-5678", "0515123456a"}
	for _, num := range invalidNumbers {
		if IsTelephone(num) {
			t.Errorf("%s should be an invalid telephone number", num)
		}
	}
}
