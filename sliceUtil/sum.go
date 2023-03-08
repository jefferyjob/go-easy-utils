package sliceUtil

func SumInt(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}
