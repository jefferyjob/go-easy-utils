package sliceUtil

func In(value interface{}, slices []interface{}) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InStr(value string, slices []string) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InInt(value int, slices []int) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InInt8(value int8, slices []int8) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InInt16(value int16, slices []int16) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InInt32(value int32, slices []int32) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InInt64(value int64, slices []int64) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InUint(value uint, slices []uint) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InUint8(value uint8, slices []uint8) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InUint16(value uint16, slices []uint16) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InUint32(value uint32, slices []uint32) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}

func InUint64(value uint64, slices []uint64) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}
