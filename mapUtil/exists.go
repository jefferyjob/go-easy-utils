package mapUtil

func MapValueExists(m map[string]interface{}, value interface{}) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

func MapKeyExists(m map[string]interface{}, key string) bool {
	_, exists := m[key]
	return exists
}
