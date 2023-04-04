package sliceUtil

// UniqueSlice 移除slice中重复的值
func UniqueSlice[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	m := make(map[T]bool)
	var result []T
	for _, v := range slice {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}
