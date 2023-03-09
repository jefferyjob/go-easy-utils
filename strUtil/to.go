package strUtil

import (
	"log"
	"strconv"
	"unsafe"
)

func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("StrToInt(%s) err:%s \n", str, err.Error())
		return 0
	}
	return i
}

func StrToInt8(str string) int8 {
	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		log.Printf("StrToInt8(%s) err:%s \n", str, err.Error())
		return 0
	}
	return int8(i)
}

func StrToInt16(str string) int16 {
	i, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		log.Printf("StrToInt16(%s) err:%s \n", str, err.Error())
		return 0
	}
	return int16(i)
}

func StrToInt32(str string) int32 {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		log.Printf("StrToInt32(%s) err:%s \n", str, err.Error())
		return 0
	}
	return int32(i)
}

func StrToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Printf("StrToInt64(%s) err:%s \n", str, err.Error())
		return 0
	}
	return i
}

// StrToBytes 字符串转字节数组
func StrToBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}
