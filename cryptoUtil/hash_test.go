package cryptoUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashSHA256(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"字符串", "hello world", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{"空字符串", "", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"数字", "1234567890", "c775e7b757ede630cd0aa1113bd102661ab38829ca52a6422ab782862f268646"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := HashSHA256(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
