package sliceUtil

// MergeSlice 将多个slice合并成一个slice
func MergeSlice[T any](slices ...[]T) []T {
	var newSlice []T
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
