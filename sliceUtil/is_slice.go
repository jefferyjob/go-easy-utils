package sliceUtil

import "reflect"

// IsSlice 判断指定值i是否是slice类型
func IsSlice(slice any) bool {
	return reflect.ValueOf(slice).Kind() == reflect.Slice
}
