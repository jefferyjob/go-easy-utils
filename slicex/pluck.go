package slicex

// Pluck 字段提取
func Pluck[T any, K comparable](items []T, keyFunc func(T) K) []K {
	var ids []K
	for _, v := range items {
		ids = append(ids, keyFunc(v))
	}
	return ids
}
