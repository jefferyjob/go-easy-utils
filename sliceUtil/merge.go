package sliceUtil

// Merge 将多个slice合并成一个slice
func Merge[T any](slices ...[]T) []T {
	var newSlice []T
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
