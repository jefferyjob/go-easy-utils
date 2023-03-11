package sliceUtil

func Merge(slices ...[]interface{}) []interface{} {
	var newSlice []interface{}
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeStr(slices ...[]string) []string {
	var newSlice []string
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeInt(slices ...[]int) []int {
	var newSlice []int
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeInt8(slices ...[]int8) []int8 {
	var newSlice []int8
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeInt16(slices ...[]int16) []int16 {
	var newSlice []int16
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeInt32(slices ...[]int32) []int32 {
	var newSlice []int32
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeInt64(slices ...[]int64) []int64 {
	var newSlice []int64
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeUint(slices ...[]uint) []uint {
	var newSlice []uint
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeUint8(slices ...[]uint8) []uint8 {
	var newSlice []uint8
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeUint16(slices ...[]uint16) []uint16 {
	var newSlice []uint16
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeUint32(slices ...[]uint32) []uint32 {
	var newSlice []uint32
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func MergeUint64(slices ...[]uint64) []uint64 {
	var newSlice []uint64
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
