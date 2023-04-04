package sliceUtil

// InSlice 判断value是否在slice中
func InSlice[T comparable](value T, slices []T) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}
