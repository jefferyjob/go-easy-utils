package mapUtil

// MapValueExists 判断map中的value是否存在
func MapValueExists[T comparable](m map[string]T, value T) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// MapKeyExists 判断map中的key是否存在
func MapKeyExists(m map[string]any, key string) bool {
	_, exists := m[key]
	return exists
}
