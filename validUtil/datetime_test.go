package validUtil

import (
	"testing"
)

func TestIsTime(t *testing.T) {
	if !IsTime("12:34:56") {
		t.Error("Expected true, got false")
	}
	if IsTime("12:34") {
		t.Error("Expected false, got true")
	}
	if IsTime("25:34:56") {
		t.Error("Expected false, got true")
	}
	if IsTime("12:60:56") {
		t.Error("Expected false, got true")
	}
	if IsTime("12:34:60") {
		t.Error("Expected false, got true")
	}
}

func TestIsDate(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"2021-01-01", true},
		{"2021/01/01", false},
		{"2021-13-01", false},
		{"2021-02-29", false},
		{"2020-02-29", true},
	}

	for _, testCase := range testCases {
		result := IsDate(testCase.input)
		if result != testCase.expected {
			t.Errorf("IsDate(%v) = %v, expected %v", testCase.input, result, testCase.expected)
		}
	}
}

func TestIsDateTime(t *testing.T) {
	if !IsDateTime("2023-03-11 12:34:56") {
		t.Error("Expected true, got false")
	}
	if IsDateTime("2023-02-29 12:34:56") {
		t.Error("Expected false, got true")
	}
	if IsDateTime("2023-04-31 12:34:56") {
		t.Error("Expected false, got true")
	}
	if IsDateTime("2023-13-01 12:34:56") {
		t.Error("Expected false, got true")
	}
	if IsDateTime("2023-03-11") {
		t.Error("Expected false, got true")
	}
	if IsDateTime("12:34:56") {
		t.Error("Expected false, got true")
	}
}
