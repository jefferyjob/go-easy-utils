package mapx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyExists(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]string
		inputKey string
		wantRes  bool
	}{
		{
			name: "存在",
			input: map[string]string{
				"for": "jack",
				"bar": "123",
			},
			inputKey: "for",
			wantRes:  true,
		},
		{
			name: "不存在",
			input: map[string]string{
				"for": "jack",
				"bar": "123",
			},
			inputKey: "tom",
			wantRes:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := KeyExists(tc.input, tc.inputKey)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestValueExists(t *testing.T) {
	testCases := []struct {
		name       string
		input      map[string]string
		inputValue string
		wantRes    bool
	}{
		{
			name: "存在",
			input: map[string]string{
				"for": "jack",
				"bar": "123",
			},
			inputValue: "jack",
			wantRes:    true,
		},
		{
			name: "不存在",
			input: map[string]string{
				"for": "jack",
				"bar": "123",
			},
			inputValue: "tom",
			wantRes:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ValueExists(tc.input, tc.inputValue)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
