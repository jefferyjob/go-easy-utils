package mathUtil

// Min 返回slice中最小值
func Min[T Numeric](slice []T) T {
	if len(slice) == 0 {
		return 0
	}
	min := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min
}
