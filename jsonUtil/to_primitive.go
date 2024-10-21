package jsonUtil

import (
	"fmt"
	"reflect"
	"strconv"
)

// 将任意数据类型强转为指定的数据类型
//
// 支持的方法如下:
//  - toStringReflect: 将任意类型的数据转为 string
//  - toInt64Reflect: 将任意类型的数据转为 int64
//  - toUint64Reflect: 将任意类型的数据转为 uint64
//  - toFloat64Reflect: 将任意类型的数据转为 float64
//  - toBoolReflect: 将任意类型的数据转为 bool
//
// 原则上更具用户定义的目标类型进行强制类型转换，通过反射机制处理不同的输入类型。
// 不同类型的处理规则如下：
//  - 指针类型：如果指针为空，则返回对应类型的零值；如果非空，则获取指针指向的值进行类型转换。
//  - 基础类型：根据反射获取的具体类型，使用相应的转换方式，例如 string 类型直接返回，int 类型转为字符串或浮点型等。
//  - 不支持类型：如果输入类型无法识别或不在支持范围内，返回该类型的零值或错误信息。

func toStringReflect(i any) string {
	if i == nil {
		return ""
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Complex64:
		return fmt.Sprintf("(%g+%gi)", real(v.Complex()), imag(v.Complex()))
	case reflect.Complex128:
		return fmt.Sprintf("(%g+%gi)", real(v.Complex()), imag(v.Complex()))
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	default:
		return ""
	}
}

func toInt64Reflect(i any) (int64, error) {
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
	case reflect.Float32, reflect.Float64:
		return int64(v.Float()), nil
	case reflect.String:
		if v.String() == "" {
			return 0, nil
		}
		intValue, err := strconv.ParseInt(v.String(), 10, 64)
		if err != nil {
			return 0, ErrSyntax
		}
		return intValue, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int64(v.Uint()), nil
	case reflect.Complex64:
		return int64(real(v.Complex())), nil
	case reflect.Complex128:
		return int64(real(v.Complex())), nil
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

func toUint64Reflect(i any) (uint64, error) {
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
			return 0, nil
		}
		return uint64(intValue), nil
	case reflect.Float32, reflect.Float64:
		floatValue := v.Float()
		if floatValue < 0 {
			return 0, nil
		}
		return uint64(floatValue), nil
	case reflect.Complex64, reflect.Complex128:
		realValue := real(v.Complex())
		if realValue < 0 {
			return 0, nil
		}
		return uint64(realValue), nil
	case reflect.String:
		if v.String() == "" {
			return 0, nil
		}
		uintValue, err := strconv.ParseUint(v.String(), 10, 64)
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

func toFloat64Reflect(i any) (float64, error) {
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
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	case reflect.String:
		if v.String() == "" {
			return 0, nil
		}
		floatValue, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return 0, ErrSyntax
		}
		return floatValue, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), nil
	case reflect.Complex64, reflect.Complex128:
		return real(v.Complex()), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	//case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice:
	//	return 0, nil
	default:
		return 0, ErrType
	}
}

func toBoolReflect(i any) bool {
	if i == nil {
		return false
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return false
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() != 0
	case reflect.String:
		val := v.String()
		if val == "true" {
			return true
		} else if val == "false" {
			return false
		}
		return val != ""
	//case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice:
	//	return !v.IsNil()
	default:
		return false
	}
}
