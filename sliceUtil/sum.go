package sliceUtil

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Sum 对slice中的元素求和
func Sum[T Numeric](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
