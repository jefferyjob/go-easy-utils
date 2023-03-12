package validUtil

import "testing"

func TestIsEmail(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"test@example.com", true},
		{"abc@def.com", true},
		{"123@456.com", true},
		{"test@.com", false},
		{"test@com", false},
		{"test@example", false},
	}

	for _, tt := range tests {
		got := IsEmail(tt.input)
		if got != tt.want {
			t.Errorf("IsEmail(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsJSON(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid json", `{"name":"John","age":30,"city":"New York"}`, true},
		{"invalid json", `{"name":"John","age":30,"city":"New York"`, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := IsJSON(tc.input); result != tc.expected {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestIsQQ(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"test1", "123456", true},
		{"test2", "12345678", true},
		{"test3", "0123456", false},
		{"test4", "10000", true},
		{"test5", "10000000000", true},
		{"test6", "9999999999", true},
		{"test7", "1", false},
		{"test8", "a1234567", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if output := IsQQ(test.input); output != test.expected {
				t.Errorf("Input: %s, Expected: %v, Output: %v", test.input, test.expected, output)
			}
		})
	}
}

func TestIsWeChat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"test1", "12345678", false},
		{"test2", "a12345678", true},
		{"test3", "abcdefghijk12345678", true},
		{"test4", "1234abcd", false},
		{"test5", "a123456789012345678", true},
		{"test6", "a_1234567", true},
		{"test7", "a-b-c-d-e-f-g-h-i-j", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if output := IsWeChat(test.input); output != test.expected {
				t.Errorf("Input: %s, Expected: %v, Output: %v", test.input, test.expected, output)
			}
		})
	}
}

func TestIsWeibo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"test1", "12345678", false},
		{"test2", "a12345678", true},
		{"test3", "abcdefghijk12345678", true},
		{"test4", "1234abcd", false},
		{"test5", "a123456789012345678", true},
		{"test6", "a_1234567", true},
		{"test7", "a-b-c-d-e-f-g-h-i-j", true},
		{"test8", "a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_r_s_t_u_v_w_x_y_z", false},
		{"test9", "a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_r_s_t_u_v_w_x_y_z_", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if output := IsWeibo(test.input); output != test.expected {
				t.Errorf("Input: %s, Expected: %v, Output: %v", test.input, test.expected, output)
			}
		})
	}
}

func TestIsPassword(t *testing.T) {
	testCases := []struct {
		password string
		expected bool
	}{
		{"", false},
		{"aBc12", false},
		{"abcdef", false},
		{"123456", false},
		{"ab@1", false},
		{"abcdefg", false},
		{"ab1", false},
		{"aBc12#", true},
		{"Abc123!@#", true},
		{"1234567Abcdefghijk#", true},
		{"_abc_123", true},
		{"_abc_123!", true},
		{"!@#$%^&*()", false},
		{"!@#$%^&*()Aa123", true},
		{"12345678901234567890", false},
		{"1234567890123456789A", false},
	}

	for _, tc := range testCases {
		result := IsPassword(tc.password)
		if result != tc.expected {
			t.Errorf("Expected %t for password '%s', but got %t", tc.expected, tc.password, result)
		}
	}
}
