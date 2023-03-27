package validUtil

import (
	"fmt"
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

func TestIsDateDemo(t *testing.T) {
	fmt.Println(IsDate("2022-04-01"))
}

func TestIsDate(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"2021-01-01", true},
		{"2021/01/01", false},
		{"2021-20-01", false},
		{"2021-02-29", false},
		{"2020-02-29", true},
		{"2020-02-25", true},
		{"2023-04-30", true},
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
