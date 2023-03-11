package sliceUtil

func SumInt(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func SumInt8(slice []int8) int8 {
	var sum int8
	for _, v := range slice {
		sum += v
	}
	return sum
}

func SumInt16(slice []int16) int16 {
	var sum int16
	for _, v := range slice {
		sum += v
	}
	return sum
}

func SumInt32(slice []int32) int32 {
	var sum int32
	for _, v := range slice {
		sum += v
	}
	return sum
}

func SumInt64(slice []int64) int64 {
	var sum int64
	for _, v := range slice {
		sum += v
	}
	return sum
}
