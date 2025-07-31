package slicex

// Filter 返回符合 f 的元素组成的新切片
func Filter[T any](s []T, f func(T) bool) []T {
	// var result []T
	result := make([]T, 0) // 保证不是 nil
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
