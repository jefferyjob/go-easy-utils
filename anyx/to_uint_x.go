package anyx

import (
	"math"
	"reflect"
	"strconv"
)

// ToUint 将给定的值转换为 uint
func ToUint(i any) (uint, error) {
	v, err := ToUint64(i)
	if err != nil {
		return 0, err
	}

	// uint 兼容32位和64位系统
	// if uint64(uint(v)) != v {
	//	return 0, ErrValOut
	// }

	return uint(v), nil
}

// ToUint8 将给定的值转换为 uint8
func ToUint8(i any) (uint8, error) {
	value, err := ToUint64(i)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint8 {
		return 0, ErrValOut
	}
	return uint8(value), nil
}

// ToUint16 将给定的值转换为 uint16
func ToUint16(i any) (uint16, error) {
	value, err := ToUint64(i)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint16 {
		return 0, ErrValOut
	}
	return uint16(value), nil
}

// ToUint32 将给定的值转换为 uint32
func ToUint32(i any) (uint32, error) {
	value, err := ToUint64(i)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint32 {
		return 0, ErrValOut
	}
	return uint32(value), nil
}

// ToUint64 将给定的值转换为 uint64
func ToUint64(i any) (uint64, error) {
	if i == nil {
		return 0, nil
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return 0, nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue := v.Int()
		if intValue < 0 {
			return 0, ErrUnsignedInt
		}
		return uint64(intValue), nil
	case reflect.Float32, reflect.Float64:
		floatValue := v.Float()
		if floatValue < 0 {
			return 0, ErrUnsignedInt
		}
		return uint64(floatValue), nil
	case reflect.Complex64, reflect.Complex128:
		realValue := real(v.Complex())
		if realValue < 0 {
			return 0, ErrUnsignedInt
		}
		return uint64(realValue), nil
	case reflect.String:
		strValue := v.String()
		uintValue, err := strconv.ParseUint(strValue, 10, 64)
		if err != nil {
			return 0, ErrSyntax
		}
		return uintValue, nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, ErrType
	}
}
