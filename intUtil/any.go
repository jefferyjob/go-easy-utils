package intUtil

import (
	"fmt"
	"strconv"
)

// AnyToInt converts any input value to int
func AnyToInt(input interface{}) (int, error) {
	switch input.(type) {
	case int:
		return input.(int), nil
	case int8:
		return int(input.(int8)), nil
	case int16:
		return int(input.(int16)), nil
	case int32:
		return int(input.(int32)), nil
	case int64:
		return int(input.(int64)), nil
	case uint:
		return int(input.(uint)), nil
	case uint8:
		return int(input.(uint8)), nil
	case uint16:
		return int(input.(uint16)), nil
	case uint32:
		return int(input.(uint32)), nil
	case uint64:
		return int(input.(uint64)), nil
	case string:
		return strconv.Atoi(input.(string))
	default:
		return 0, fmt.Errorf("unsupported type %T", input)
	}
}

// AnyToInt8 converts any input value to int8
func AnyToInt8(input interface{}) (int8, error) {
	value, err := AnyToInt(input)
	if err != nil {
		return 0, err
	}
	if value < -128 || value > 127 {
		return 0, fmt.Errorf("%d out of int8 range", value)
	}
	return int8(value), nil
}

// AnyToInt16 converts any input value to int16
func AnyToInt16(input interface{}) (int16, error) {
	value, err := AnyToInt(input)
	if err != nil {
		return 0, err
	}
	if value < -32768 || value > 32767 {
		return 0, fmt.Errorf("%d out of int16 range", value)
	}
	return int16(value), nil
}

// AnyToInt32 converts any input value to int32
func AnyToInt32(input interface{}) (int32, error) {
	value, err := AnyToInt(input)
	if err != nil {
		return 0, err
	}
	if value < -2147483648 || value > 2147483647 {
		return 0, fmt.Errorf("%d out of int32 range", value)
	}
	return int32(value), nil
}

// AnyToInt64 converts any input value to int64
func AnyToInt64(input interface{}) (int64, error) {
	value, err := AnyToInt(input)
	if err != nil {
		return 0, err
	}
	return int64(value), nil
}

// AnyToUint converts any input value to uint
func AnyToUint(input interface{}) (uint, error) {
	value, err := AnyToInt(input)
	if err != nil {
		return 0, err
	}
	if value < 0 {
		return 0, fmt.Errorf("%d is negative", value)
	}
	return uint(value), nil
}

// AnyToUint8 converts any input value to uint8
func AnyToUint8(input interface{}) (uint8, error) {
	value, err := AnyToUint(input)
	if err != nil {
		return 0, err
	}
	if value > 255 {
		return 0, fmt.Errorf("%d out of uint8 range", value)
	}
	return uint8(value), nil
}

// AnyToUint16 converts any input value to uint16
func AnyToUint16(input interface{}) (uint16, error) {
	value, err := AnyToUint(input)
	if err != nil {
		return 0, err
	}
	if value > 65535 {
		return 0, fmt.Errorf("%d out of uint16 range", value)
	}
	return uint16(value), nil
}

// AnyToUint32 converts any input value to uint32
func AnyToUint32(input interface{}) (uint32, error) {
	value, err := AnyToUint(input)
	if err != nil {
		return 0, err
	}
	if value > 4294967295 {
		return 0, fmt.Errorf("%d out of uint32 range", value)
	}
	return uint32(value), nil
}

// AnyToUint64 converts any input value to uint64
func AnyToUint64(input interface{}) (uint64, error) {
	value, err := AnyToUint(input)
	if err != nil {
		return 0, err
	}
	return uint64(value), nil
}
