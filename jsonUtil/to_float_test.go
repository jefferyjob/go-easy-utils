package jsonUtil

import (
	go_easy_utils "github.com/jefferyjob/go-easy-utils"
	"reflect"
	"testing"
)

//func TestFloat(t *testing.T) {
//	var data1 interface{}
//	data1 = 123
//	fmt.Println(toFloat64Reflect(data1))
//
//	data2 := int16Ptr(80)
//	fmt.Println(toFloat64Reflect(data2))
//
//	data3 := map[interface{}]interface{}{
//		"aaa": "aaa",
//	}
//	fmt.Println(toFloat64Reflect(data3))
//}

func TestToFloat64(t *testing.T) {
	testCases := []struct {
		value       interface{}
		expected    float64
		expectedErr error
	}{
		{nil, 0, nil},
		{float32(123.5), 123.5, nil},
		{"123.456", 123.456, nil},
		{uint(123), 123, nil},
		{uint8(123), 123, nil},
		{uint16(123), 123, nil},
		{uint32(123), 123, nil},
		{uint64(123), 123, nil},
		{int(123), 123, nil},
		{int8(123), 123, nil},
		{int16(123), 123, nil},
		{int32(123), 123, nil},
		{int64(123), 123, nil},
		{complex64(1 + 2i), 1, nil},
		{complex128(1 + 2i), 1, nil},
		{true, 1, nil},
		{false, 0, nil},
		{(*bool)(nil), 0, nil},
		{make(chan int), 0, go_easy_utils.ErrType},
		{"abc", 0, go_easy_utils.ErrSyntax},
	}

	for _, tc := range testCases {
		f, err := toFloat64(tc.value)
		if f != tc.expected || err != tc.expectedErr {
			t.Errorf("toFloat64(%v) = (%v, %v), expected (%v, %v)", tc.value, f, err, tc.expected, tc.expectedErr)
		}

		f2, err2 := toFloat64Reflect(tc.value)
		if f != tc.expected || err2 != tc.expectedErr {
			t.Errorf("toFloat64Reflect(%v) = (%v, %v), expected (%v, %v)", tc.value, f2, err2, tc.expected, tc.expectedErr)
		}
	}
}

func TestToFloat64Pointer(t *testing.T) {
	testCases := []struct {
		name          string
		input         interface{}
		expectedValue float64
		expectedError error
	}{
		{
			name:          "nil input",
			input:         nil,
			expectedValue: 0,
			expectedError: nil,
		},
		{
			name:          "float32 pointer input",
			input:         float32Ptr(12.2),
			expectedValue: 12.199999809265137,
			expectedError: nil,
		},
		{
			name:          "float64 pointer input",
			input:         float64Ptr(56.78),
			expectedValue: 56.78,
			expectedError: nil,
		},
		{
			name:          "invalid string input",
			input:         stringPtr("abc"),
			expectedValue: 0,
			expectedError: go_easy_utils.ErrSyntax,
		},
		{
			name:          "valid string input",
			input:         stringPtr("123.45"),
			expectedValue: 123.45,
			expectedError: nil,
		},
		{
			name:          "uint pointer input",
			input:         uintPtr(10),
			expectedValue: 10,
			expectedError: nil,
		},
		{
			name:          "uint8 pointer input",
			input:         uint8Ptr(20),
			expectedValue: 20,
			expectedError: nil,
		},
		{
			name:          "uint16 pointer input",
			input:         uint16Ptr(30),
			expectedValue: 30,
			expectedError: nil,
		},
		{
			name:          "uint32 pointer input",
			input:         uint32Ptr(40),
			expectedValue: 40,
			expectedError: nil,
		},
		{
			name:          "uint64 pointer input",
			input:         uint64Ptr(50),
			expectedValue: 50,
			expectedError: nil,
		},
		{
			name:          "int pointer input",
			input:         intPtr(-60),
			expectedValue: -60,
			expectedError: nil,
		},
		{
			name:          "int8 pointer input",
			input:         int8Ptr(-70),
			expectedValue: -70,
			expectedError: nil,
		},
		{
			name:          "int16 pointer input",
			input:         int16Ptr(-80),
			expectedValue: -80,
			expectedError: nil,
		},
		{
			name:          "int32 pointer input",
			input:         int32Ptr(-90),
			expectedValue: -90,
			expectedError: nil,
		},
		{
			name:          "int64 pointer input",
			input:         int64Ptr(-100),
			expectedValue: -100,
			expectedError: nil,
		},
		{
			name:          "complex64 pointer input",
			input:         complex64Ptr(complex(1, 2)),
			expectedValue: 1,
			expectedError: nil,
		},
		{
			name:          "complex128 pointer input",
			input:         complex128Ptr(complex(3, 4)),
			expectedValue: 3,
			expectedError: nil,
		},
		{
			name:          "bool pointer input - true",
			input:         boolPtr(true),
			expectedValue: 1,
			expectedError: nil,
		},
		{
			name:          "bool pointer input - false",
			input:         boolPtr(false),
			expectedValue: 0,
			expectedError: nil,
		},
		{
			name:          "unsupported input type",
			input:         "unsupported",
			expectedValue: 0,
			expectedError: go_easy_utils.ErrSyntax,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualValue, actualError := toFloat64(tc.input)

			if actualError != tc.expectedError {
				t.Errorf("Expected(name:%s) error: %v, but got: %v", tc.name, tc.expectedError, actualError)
			}

			if !reflect.DeepEqual(actualValue, tc.expectedValue) {
				t.Errorf("Expected value: %v, but got: %v", tc.expectedValue, actualValue)
			}

			// Reflect
			actualValue2, actualError2 := toFloat64Reflect(tc.input)

			if actualError2 != tc.expectedError {
				t.Errorf("Expected(name:%s) error: %v, but got: %v", tc.name, tc.expectedError, actualError2)
			}

			if !reflect.DeepEqual(actualValue2, tc.expectedValue) {
				t.Errorf("Expected value: %v, but got: %v", tc.expectedValue, actualValue2)
			}
		})
	}
}

func float32Ptr(v float32) *float32 {
	return &v
}

func float64Ptr(v float64) *float64 {
	return &v
}

func stringPtr(v string) *string {
	return &v
}

func uintPtr(v uint) *uint {
	return &v
}

func uint8Ptr(v uint8) *uint8 {
	return &v
}

func uint16Ptr(v uint16) *uint16 {
	return &v
}

func uint32Ptr(v uint32) *uint32 {
	return &v
}

func uint64Ptr(v uint64) *uint64 {
	return &v
}

func intPtr(v int) *int {
	return &v
}

func int8Ptr(v int8) *int8 {
	return &v
}

func int16Ptr(v int16) *int16 {
	return &v
}

func int32Ptr(v int32) *int32 {
	return &v
}

func int64Ptr(v int64) *int64 {
	return &v
}

func complex64Ptr(v complex64) *complex64 {
	return &v
}

func complex128Ptr(v complex128) *complex128 {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}
