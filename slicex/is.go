package slicex

import "reflect"

// Is 判断指定值i是否是slice类型
func Is(slice any) bool {
	return reflect.ValueOf(slice).Kind() == reflect.Slice
}
