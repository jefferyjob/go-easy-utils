package cryptoUtil

import (
	"testing"
)

func TestMd5(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "",
			output: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			input:  "test",
			output: "098f6bcd4621d373cade4e832627b4f6",
		},
		{
			input:  "hello world",
			output: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
	}

	for _, tc := range testCases {
		result := Md5(tc.input)
		if result != tc.output {
			t.Errorf("Md5(%q) = %q; want %q", tc.input, result, tc.output)
		}
	}
}
