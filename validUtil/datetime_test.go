package validUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDate(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"合法日期", "2021-01-01", true},
		{"非法分隔符", "2021/01/01", false},
		{"非法月份", "2021-20-01", false},
		{"非闰年2月29", "2021-02-29", false},
		{"闰年2月29", "2020-02-29", true},
		{"合法2月25", "2020-02-25", true},
		{"合法4月30", "2023-04-30", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsDate(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestIsDateTime(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"合法日期时间", "2023-03-11 12:34:56", true},
		{"无效闰日", "2023-02-29 12:34:56", false},
		{"无效日期", "2023-04-31 12:34:56", false},
		{"无效月份", "2023-13-01 12:34:56", false},
		{"仅日期", "2023-03-11", false},
		{"仅时间", "12:34:56", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IsDateTime(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestIsValidDay(t *testing.T) {
	testCases := []struct {
		name  string
		day   int
		month int
		year  int
		want  bool
	}{
		{"Valid day in Jan", 31, 1, 2024, true},
		{"Invalid day in Jan", 32, 1, 2024, false},
		{"Valid day in Apr", 30, 4, 2024, true},
		{"Invalid day in Apr", 31, 4, 2024, false},
		{"Valid day in Feb (Leap year)", 29, 2, 2024, true},
		{"Invalid day in Feb (Leap year)", 30, 2, 2024, false},
		{"Valid day in Feb (Non-leap year)", 28, 2, 2023, true},
		{"Invalid day in Feb (Non-leap year)", 29, 2, 2023, false},
		{"Valid day in Dec", 31, 12, 2024, true},
		{"Invalid day in Dec", 32, 12, 2024, false},
		{"Valid day at edge case", 1, 1, 2024, true},
		{"Invalid day at edge case", 0, 1, 2024, false},
		{"Invalid month", 15, 13, 2024, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isValidDay(tc.day, tc.month, tc.year)
			assert.Equal(t, tc.want, res)
		})
	}
}
