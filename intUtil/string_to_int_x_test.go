package intUtil

import (
	"testing"
)

func TestStrToInt(t *testing.T) {
	s := "123"
	expected := 123
	result, err := StrToInt(s)
	if err != nil {
		t.Errorf("StrToInt(%s) returned an error: %v", s, err)
	}
	if result != expected {
		t.Errorf("StrToInt(%s) = %d; expected %d", s, result, expected)
	}
}

func TestStrToInt8(t *testing.T) {
	s := "-123"
	expected := int8(-123)
	result, err := StrToInt8(s)
	if err != nil {
		t.Errorf("StrToInt8(%s) returned an error: %v", s, err)
	}
	if result != expected {
		t.Errorf("StrToInt8(%s) = %d; expected %d", s, result, expected)
	}
}

func TestStrToInt16(t *testing.T) {
	s := "123"
	expected := int16(123)
	result, err := StrToInt16(s)
	if err != nil {
		t.Errorf("StrToInt16(%s) returned an error: %v", s, err)
	}
	if result != expected {
		t.Errorf("StrToInt16(%s) = %d; expected %d", s, result, expected)
	}
}

func TestStrToInt32(t *testing.T) {
	s := "-123"
	expected := int32(-123)
	result, err := StrToInt32(s)
	if err != nil {
		t.Errorf("StrToInt32(%s) returned an error: %v", s, err)
	}
	if result != expected {
		t.Errorf("StrToInt32(%s) = %d; expected %d", s, result, expected)
	}
}

func TestStrToInt64(t *testing.T) {
	s := "123"
	expected := int64(123)
	result, err := StrToInt64(s)
	if err != nil {
		t.Errorf("StrToInt64(%s) returned an error: %v", s, err)
	}
	if result != expected {
		t.Errorf("StrToInt64(%s) = %d; expected %d", s, result, expected)
	}
}
