package strx

import "strconv"

// ToUint string转uint
func ToUint(v string) uint {
	i, err := strconv.ParseUint(v, 10, 0)
	if err != nil {
		return 0
	}
	return uint(i)
}

// ToUint8 string转uint8
func ToUint8(v string) uint8 {
	i, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0
	}
	return uint8(i)
}

// ToUint16 string转uint16
func ToUint16(v string) uint16 {
	i, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		return 0
	}
	return uint16(i)
}

// ToUint32 string转uint32
func ToUint32(v string) uint32 {
	i, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(i)
}

// ToUint64 string转uint64
func ToUint64(v string) uint64 {
	i, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
