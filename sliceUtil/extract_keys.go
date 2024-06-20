package sliceUtil

// ExtractKeys 字段提取
func ExtractKeys[T any, K comparable](items []T, keyFunc func(T) K) []K {
	var ids []K
	for _, v := range items {
		ids = append(ids, keyFunc(v))
	}
	return ids
}
