package sliceUtil

// InSlices 判断指定值value是否在指定的slice中存在
func InSlices(value interface{}, slices []interface{}) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InStrSlices 判断指定值value是否在指定的slice中存在
func InStrSlices(value string, slices []string) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InIntSlices 判断指定值value是否在指定的slice中存在
func InIntSlices(value int, slices []int) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InInt8Slices 判断指定值value是否在指定的slice中存在
func InInt8Slices(value int8, slices []int8) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InInt16Slices 判断指定值value是否在指定的slice中存在
func InInt16Slices(value int16, slices []int16) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InInt32Slices 判断指定值value是否在指定的slice中存在
func InInt32Slices(value int32, slices []int32) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InInt64Slices 判断指定值value是否在指定的slice中存在
func InInt64Slices(value int64, slices []int64) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InUintSlices 判断指定值value是否在指定的slice中存在
func InUintSlices(value uint, slices []uint) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InUint8Slices 判断指定值value是否在指定的slice中存在
func InUint8Slices(value uint8, slices []uint8) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InUint16Slices 判断指定值value是否在指定的slice中存在
func InUint16Slices(value uint16, slices []uint16) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InUint32Slices 判断指定值value是否在指定的slice中存在
func InUint32Slices(value uint32, slices []uint32) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

// InUint64Slices 判断指定值value是否在指定的slice中存在
func InUint64Slices(value uint64, slices []uint64) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}
