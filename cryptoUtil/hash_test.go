package cryptoUtil

import "testing"

func TestHashSHA256(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello world", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"1234567890", "c775e7b757ede630cd0aa1113bd102661ab38829ca52a6422ab782862f268646"},
	}

	for _, tc := range testCases {
		output := HashSHA256(tc.input)
		if output != tc.expected {
			t.Errorf("HashSHA256(%s) = %s; expected %s", tc.input, output, tc.expected)
		}
	}
}
