package sliceUtil

// SumIntSlice 对slice中的元素求和
func SumIntSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumInt8Slice 对slice中的元素求和
func SumInt8Slice(slice []int8) int8 {
	var sum int8
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumInt16Slice 对slice中的元素求和
func SumInt16Slice(slice []int16) int16 {
	var sum int16
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumInt32Slice 对slice中的元素求和
func SumInt32Slice(slice []int32) int32 {
	var sum int32
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumInt64Slice 对slice中的元素求和
func SumInt64Slice(slice []int64) int64 {
	var sum int64
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumFloat32Slice 对slice中的元素求和
func SumFloat32Slice(slice []float32) float32 {
	var sum float32
	for _, v := range slice {
		sum += v
	}
	return sum
}

// SumFloat64Slice 对slice中的元素求和
func SumFloat64Slice(slice []float64) float64 {
	var sum float64
	for _, v := range slice {
		sum += v
	}
	return sum
}
