package mapUtil

// MapValueExists 判断map中的value是否存在
func MapValueExists(m map[string]interface{}, value interface{}) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// MapKeyExists 判断map中的key是否存在
func MapKeyExists(m map[string]interface{}, key string) bool {
	_, exists := m[key]
	return exists
}
