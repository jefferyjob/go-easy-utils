package slicex

// Merge 将多个slice合并成一个slice
func Merge[T any](slices ...[]T) []T {
	var new []T
	for _, slice := range slices {
		for _, v := range slice {
			new = append(new, v)
		}
	}
	return new
}
