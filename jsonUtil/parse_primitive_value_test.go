package jsonUtil

import (
	"reflect"
	"testing"
)

func TestParsePrimitiveValue(t *testing.T) {
	// Test cases for reflect.String
	strFieldVal := reflect.ValueOf(new(string)).Elem()
	err := parsePrimitiveValue(strFieldVal, "hello")
	if err != nil {
		t.Errorf("Error parsing string value: %v", err)
	}
	if strFieldVal.String() != "hello" {
		t.Errorf("Expected %q, got %q", "hello", strFieldVal.String())
	}
	err = parsePrimitiveValue(strFieldVal, 123)
	if err == nil {
		t.Error("Expected error for parsing int value as string")
	}

	// Test cases for reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64
	intFieldVal := reflect.ValueOf(new(int64)).Elem()
	err = parsePrimitiveValue(intFieldVal, "123")
	if err != nil {
		t.Errorf("Error parsing int value: %v", err)
	}
	if intFieldVal.Int() != 123 {
		t.Errorf("Expected %d, got %d", 123, intFieldVal.Int())
	}
	err = parsePrimitiveValue(intFieldVal, "abc")
	if err == nil {
		t.Error("Expected error for parsing non-int value as int")
	}

	// Test cases for reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64
	uintFieldVal := reflect.ValueOf(new(uint64)).Elem()
	err = parsePrimitiveValue(uintFieldVal, "123")
	if err != nil {
		t.Errorf("Error parsing uint value: %v", err)
	}
	if uintFieldVal.Uint() != 123 {
		t.Errorf("Expected %d, got %d", 123, uintFieldVal.Uint())
	}
	err = parsePrimitiveValue(uintFieldVal, 123)
	if err != nil {
		t.Errorf("Expected error  uint, err:%v", err)
	}

	// Test cases for reflect.Float32, reflect.Float64
	floatFieldVal := reflect.ValueOf(new(float64)).Elem()
	err = parsePrimitiveValue(floatFieldVal, "3.14")
	if err != nil {
		t.Errorf("Error parsing float value: %v", err)
	}
	if floatFieldVal.Float() != 3.14 {
		t.Errorf("Expected %f, got %f", 3.14, floatFieldVal.Float())
	}
	err = parsePrimitiveValue(floatFieldVal, "abc")
	if err == nil {
		t.Error("Expected error for parsing non-float value as float")
	}

	// Test cases for reflect.Bool
	boolFieldVal := reflect.ValueOf(new(bool)).Elem()
	err = parsePrimitiveValue(boolFieldVal, true)
	if err != nil {
		t.Errorf("Error parsing bool value: %v", err)
	}
	if boolFieldVal.Bool() != true {
		t.Errorf("Expected %t, got %t", true, boolFieldVal.Bool())
	}
	err = parsePrimitiveValue(boolFieldVal, "abc")
	if err == nil {
		t.Error("Expected error for parsing non-bool value as bool")
	}

	// Test cases for unsupported kind
	unsupportedFieldVal := reflect.ValueOf(new([]string)).Elem()
	err = parsePrimitiveValue(unsupportedFieldVal, "abc")
	if err == nil {
		t.Error("Expected error for unsupported kind")
	}
}
