package validUtil

import "testing"

func TestIsBankCardNo(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid bank card number with 15 digits",
			input:    "123456789012345",
			expected: false,
		},
		{
			name:     "invalid bank card number with less than 15 digits",
			input:    "123456789012",
			expected: false,
		},
		{
			name:     "invalid bank card number with more than 19 digits",
			input:    "12345678901234567890",
			expected: false,
		},
		{
			name:     "invalid bank card number with non-digit characters",
			input:    "12345abc6789012345",
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsBankCardNo(tc.input)
			if result != tc.expected {
				t.Errorf("input(%s) expected %v, but got %v", tc.input, tc.expected, result)
			}
		})
	}
}
