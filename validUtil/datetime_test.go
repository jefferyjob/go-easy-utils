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
	if !IsDate("2023-03-11") {
		t.Error("Expected true, got false")
	}
	if IsDate("2023-02-29") {
		t.Error("Expected false, got true")
	}
	if IsDate("2023-04-31") {
		t.Error("Expected false, got true")
	}
	if IsDate("2023-13-01") {
		t.Error("Expected false, got true")
	}
	if IsDate("2023-03-11 12:34:56") {
		t.Error("Expected false, got true")
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
