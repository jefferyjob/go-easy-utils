package mapUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapKeyExists(t *testing.T) {
	testCases := []struct {
		name     string
		inputMap map[string]string
		inputKey string
		wantRes  bool
	}{
		{
			name: "存在",
			inputMap: map[string]string{
				"for": "jack",
				"bar": "123",
			},
			inputKey: "for",
			wantRes:  true,
		},
		{
			name: "不存在",
			inputMap: map[string]string{
				"for": "jack",
				"bar": "123",
			},
			inputKey: "tom",
			wantRes:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := MapKeyExists(tc.inputMap, tc.inputKey)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestMapValueExists(t *testing.T) {
	testCases := []struct {
		name       string
		inputMap   map[string]string
		inputValue string
		wantRes    bool
	}{
		{
			name: "存在",
			inputMap: map[string]string{
				"for": "jack",
				"bar": "123",
			},
			inputValue: "jack",
			wantRes:    true,
		},
		{
			name: "不存在",
			inputMap: map[string]string{
				"for": "jack",
				"bar": "123",
			},
			inputValue: "tom",
			wantRes:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := MapValueExists(tc.inputMap, tc.inputValue)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
