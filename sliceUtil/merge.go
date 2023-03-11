package sliceUtil

// MergeSlices 将多个slice合并成一个slice
func MergeSlices(slices ...[]interface{}) []interface{} {
	var newSlice []interface{}
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeStrSlices 将多个slice合并成一个slice
func MergeStrSlices(slices ...[]string) []string {
	var newSlice []string
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeIntSlices 将多个slice合并成一个slice
func MergeIntSlices(slices ...[]int) []int {
	var newSlice []int
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeInt8Slices 将多个slice合并成一个slice
func MergeInt8Slices(slices ...[]int8) []int8 {
	var newSlice []int8
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeInt16Slices 将多个slice合并成一个slice
func MergeInt16Slices(slices ...[]int16) []int16 {
	var newSlice []int16
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeInt32Slices 将多个slice合并成一个slice
func MergeInt32Slices(slices ...[]int32) []int32 {
	var newSlice []int32
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeInt64Slices 将多个slice合并成一个slice
func MergeInt64Slices(slices ...[]int64) []int64 {
	var newSlice []int64
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeUintSlices 将多个slice合并成一个slice
func MergeUintSlices(slices ...[]uint) []uint {
	var newSlice []uint
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeUint8Slices 将多个slice合并成一个slice
func MergeUint8Slices(slices ...[]uint8) []uint8 {
	var newSlice []uint8
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeUint16Slices 将多个slice合并成一个slice
func MergeUint16Slices(slices ...[]uint16) []uint16 {
	var newSlice []uint16
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeUint32Slices 将多个slice合并成一个slice
func MergeUint32Slices(slices ...[]uint32) []uint32 {
	var newSlice []uint32
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// MergeUint64Slices 将多个slice合并成一个slice
func MergeUint64Slices(slices ...[]uint64) []uint64 {
	var newSlice []uint64
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
