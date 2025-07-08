package floatx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat32ToStr(t *testing.T) {
	testCases := []struct {
		name     string
		input    float32
		expected string
	}{
		{"一位小数", float32(3.1), "3.1"},
		{"多位小数", float32(3.122), "3.122"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Float32ToStr(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestFloat64ToStr(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected string
	}{
		{"一位小数", float64(3.1), "3.1"},
		{"多位小数", float64(3.122), "3.122"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Float64ToStr(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestFloat32ToFloat64(t *testing.T) {
	testCases := []struct {
		name     string
		input    float32
		expected float64
	}{
		{"一位小数", float32(3.1), float64(3.1)},
		{"多位小数", float32(3.122), float64(3.122)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Float32ToFloat64(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestFloat64ToFloat32(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected float32
	}{
		{"一位小数", float64(3.1), float32(3.1)},
		{"多位小数", float64(3.122), float32(3.122)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Float64ToFloat32(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
