package anyUtil

import "fmt"

// AnyToUint 将给定的值转换为 uint
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

// AnyToUint8 将给定的值转换为 uint8
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

// AnyToUint16 将给定的值转换为 uint16
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

// AnyToUint32 将给定的值转换为 uint32
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

// AnyToUint64 将给定的值转换为 uint64
func AnyToUint64(input interface{}) (uint64, error) {
	value, err := AnyToUint(input)
	if err != nil {
		return 0, err
	}
	return uint64(value), nil
}
