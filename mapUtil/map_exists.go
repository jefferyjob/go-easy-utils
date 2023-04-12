package mapUtil

// MapKeyExists 判断map中的key是否存在
func MapKeyExists[T comparable, T2 any](m map[T]T2, key T) bool {
	_, exists := m[key]
	return exists
}

// MapValueExists 判断map中的value是否存在
func MapValueExists[T comparable, T2 comparable](m map[T2]T, value T) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}
