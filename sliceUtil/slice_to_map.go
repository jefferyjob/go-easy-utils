package sliceUtil

// SliceToMap 切片转字典
func SliceToMap[T any, K comparable](items []T, keyFunc func(T) K) map[K]T {
	m := make(map[K]T)
	for _, item := range items {
		m[keyFunc(item)] = item
	}
	return m
}
