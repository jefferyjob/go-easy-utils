package strUtil

import (
	"testing"
)

func TestStrToInt(t *testing.T) {
	expected := 123
	result := StrToInt("123")
	if result != expected {
		t.Errorf("StrToInt test failed, expected %d but got %d", expected, result)
	}
}

func TestStrToInt8(t *testing.T) {
	expected := int8(123)
	result := StrToInt8("123")
	if result != expected {
		t.Errorf("StrToInt8 test failed, expected %d but got %d", expected, result)
	}
}

func TestStrToInt16(t *testing.T) {
	expected := int16(123)
	result := StrToInt16("123")
	if result != expected {
		t.Errorf("StrToInt16 test failed, expected %d but got %d", expected, result)
	}
}

func TestStrToInt32(t *testing.T) {
	expected := int32(123)
	result := StrToInt32("123")
	if result != expected {
		t.Errorf("StrToInt32 test failed, expected %d but got %d", expected, result)
	}
}

func TestStrToInt64(t *testing.T) {
	expected := int64(123)
	result := StrToInt64("123")
	if result != expected {
		t.Errorf("StrToInt64 test failed, expected %d but got %d", expected, result)
	}
}
