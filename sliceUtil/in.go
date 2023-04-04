package sliceUtil

// In 判断指定值value是否在指定的slice中存在
func In[T comparable](value T, slices []T) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}
