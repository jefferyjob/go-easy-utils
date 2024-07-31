package validUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"合法邮箱", "test@example.com", true},
		{"合法邮箱", "abc@def.com", true},
		{"合法邮箱", "123@456.com", true},
		{"非法邮箱", "test@.com", false},
		{"非法邮箱", "test@com", false},
		{"非法邮箱", "test@example", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := IsEmail(tc.input)
			assert.Equal(t, tc.want, res)
		})
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
			res := IsJSON(tc.input)
			assert.Equal(t, tc.expected, res)
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := IsQQ(tc.input)
			assert.Equal(t, tc.expected, res)
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := IsWeChat(tc.input)
			assert.Equal(t, tc.expected, res)
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := IsWeibo(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestIsPassword(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"空密码", "", false},
		{"短混合", "aBc12", false},
		{"全小写", "abcdef", false},
		{"全数字", "123456", false},
		{"短特殊", "ab@1", false},
		{"仅字母", "abcdefg", false},
		{"仅数字", "ab1", false},
		{"有效短", "aBc12#", true},
		{"有效长", "Abc123!@#", true},
		{"长有效", "1234567Abcdefghijk#", true},
		{"下划线数", "_abc_123", true},
		{"下划线特", "_abc_123!", true},
		{"全特殊", "!@#$%^&*()", false},
		{"特大小写", "!@#$%^&*()Aa123", true},
		{"超长无大", "12345678901234567890", false},
		{"超长混合", "1234567890123456789A", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsPassword(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
