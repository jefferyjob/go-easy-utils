# go-easy-utils

[![GoEasyUtils Release](https://img.shields.io/github/release/jefferyjob/go-easy-utils.svg)](https://github.com/jefferyjob/go-easy-utils/releases)
[![Go Action](https://github.com/jefferyjob/go-easy-utils/workflows/Go/badge.svg?branch=master)](https://github.com/jefferyjob/go-easy-utils/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/jefferyjob/go-easy-utils)](https://goreportcard.com/report/github.com/jefferyjob/go-easy-utils)
[![Go Coverage](https://codecov.io/gh/jefferyjob/go-easy-utils/branch/master/graph/badge.svg)](https://codecov.io/gh/jefferyjob/go-easy-utils)
[![GoEasyUtils Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/jefferyjob/go-easy-utils)
![License](https://img.shields.io/github/license/jefferyjob/go-easy-utils)

[English](README.md) | 简体中文

## 介绍

这是一个基于 Go 语言开发的通用数据类型处理工具类，帮助开发者在业务代码实现中处理常见的数据类型和数据操作。可以让您专注于您的业务代码的实现，而免去处理基本数据类型转换和验证的功能。该工具库无侵入式的设计可以让您的业务代码更容易阅读和优雅。

## 快速开始
**安装**
```bash
go get -u github.com/jefferyjob/go-easy-utils
```
**使用Demo**
```go
package main

import (
	"fmt"
	"github.com/jefferyjob/go-easy-utils/sliceUtil"
)

func main() {
	var slice = []string{"this", "is", "go", "easy", "utils"}
	chunkSlice := sliceUtil.ChunkStr(slice, 2)
	fmt.Printf("%v", chunkSlice)
}
```

## 功能列表

### jsonUtil Json处理工具

```go
// JsonToStruct 将 JSON 字符串解析为指定的结构体指针
func JsonToStruct(jsonData string, result interface{}) error
```

### ValidUtil 验证工具

```go
// IsTime 验证是否为时间格式（HH:mm:ss）
func IsTime(str string) bool

// IsDate 验证是否为日期格式（yyyy-MM-dd）
func IsDate(str string) bool

// IsDateTime 验证是否为日期时间格式（yyyy-MM-dd HH:mm:ss）
func IsDateTime(str string) bool

// IsIDCard 验证身份证号(18或15位)
func IsIDCard(str string) bool

// IsIDCard18 验证18位身份证号
func IsIDCard18(id string) bool

// IsIDCard15 验证15位身份证号
func IsIDCard15(idCard string) bool

// IsMobile 验证是否为手机号码
func IsMobile(mobileNum string) bool 

// IsTelephone 验证是否为座机号码
func IsTelephone(telephone string) bool 

// IsPostalCode 验证是否为邮编号码
func IsPostalCode(str string) bool 

// IsDecimal 验证给定的字符串小数点后是否最多两位
func IsDecimal(input string) bool 

// IsNumber 验证是否全部为数字
func IsNumber(input string) bool

// IsBankCardNo 验证是否为银行卡号
func IsBankCardNo(str string) bool

// IsAllChinese 验证给定的字符串全部为中文
func IsAllChinese(input string) bool

// IsContainChinese 验证给定的字符串包含中文
func IsContainChinese(input string) bool

// IsEmail 是否为email
func IsEmail(input string) bool

// IsIPv4 是否为ipv4地址
func IsIPv4(input string) bool

// IsIPv6 是否为ipv6地址
func IsIPv6(input string) bool

// IsURL 是否为URL地址
func IsURL(input string) bool

// IsJSON 是否为Json
func IsJSON(input string) bool

// IsChineseName 验证是否为中文名
func IsChineseName(name string) bool

// IsEnglishName 验证是否为英文名
func IsEnglishName(name string) bool

// IsQQ 验证是否为QQ号
func IsQQ(qq string) bool 

// IsWeChat 验证是否为微信号
func IsWeChat(wechat string) bool

// IsWeibo 验证是否为微博ID
func IsWeibo(weibo string) bool

// IsPassword 验证密码是否合法
// 密码长度在6-20个字符之间，只包含数字、字母和下划线
func IsPassword(password string) bool
```

### strUtil 字符串工具

```go
// StrToInt string转int
func StrToInt(str string) int 

// StrToInt8 string转int8
func StrToInt8(str string) int8 

// StrToInt16 string转int16
func StrToInt16(str string) int16 

// StrToInt32 string转int32
func StrToInt32(str string) int32 

// StrToInt64 string转int64
func StrToInt64(str string) int64 

// StrToBytes 字符串转字节数组
func StrToBytes(data string) []byte 
```

### sliceUtil 切片处理工具

```go
// Chunk 将一个切片按指定大小分成多个切片，并返回一个包含这些切片的二维切片。
func Chunk(slice []interface{}, size int) [][]interface{} 

// ChunkStr 将一个 string 类型的切片切割成指定大小的多个子切片
func ChunkStr(slice []string, size int) [][]string

// ChunkInt 将一个 int 类型的切片切割成指定大小的多个子切片
func ChunkInt(slice []int, size int) [][]int 

// ChunkInt8 将一个 int8 类型的切片切割成指定大小的多个子切片
func ChunkInt8(slice []int8, size int) [][]int8

// ChunkInt32 将一个 int32 类型的切片切割成指定大小的多个子切片
func ChunkInt32(slice []int32, size int) [][]int32

// ChunkInt64 将一个 int64 类型的切片切割成指定大小的多个子切片
func ChunkInt64(slice []int64, size int) [][]int64

// ChunkUint 将一个 uint 类型的切片切割成指定大小的多个子切片
func ChunkUint(slice []uint, size int) [][]uint

// ChunkUint8 将一个 uint8 类型的切片切割成指定大小的多个子切片
func ChunkUint8(slice []uint8, size int) [][]uint8

// ChunkUint16 将一个 uint16 类型的切片切割成指定大小的多个子切片
func ChunkUint16(slice []uint16, size int) [][]uint16

// ChunkUint32 将一个 uint32 类型的切片切分成固定大小的子切片
func ChunkUint32(slice []uint32, size int) [][]uint32

// ChunkUint64 将一个 uint64 类型的切片切分成固定大小的子切片
func ChunkUint64(slice []uint64, size int) [][]uint64

// InSlices 判断指定值value是否在指定的slice中存在
func InSlices(value interface{}, slices []interface{}) bool

// InStrSlices 判断指定值value是否在指定的slice中存在
func InStrSlices(value string, slices []string) bool 

// InIntSlices 判断指定值value是否在指定的slice中存在
func InIntSlices(value int, slices []int) bool

// InInt8Slices 判断指定值value是否在指定的slice中存在
func InInt8Slices(value int8, slices []int8) bool

// InInt16Slices 判断指定值value是否在指定的slice中存在
func InInt16Slices(value int16, slices []int16) bool 

// InInt32Slices 判断指定值value是否在指定的slice中存在
func InInt32Slices(value int32, slices []int32) bool

// InInt64Slices 判断指定值value是否在指定的slice中存在
func InInt64Slices(value int64, slices []int64) bool 

// InUintSlices 判断指定值value是否在指定的slice中存在
func InUintSlices(value uint, slices []uint) bool

// InUint8Slices 判断指定值value是否在指定的slice中存在
func InUint8Slices(value uint8, slices []uint8) bool

// InUint16Slices 判断指定值value是否在指定的slice中存在
func InUint16Slices(value uint16, slices []uint16) bool

// InUint32Slices 判断指定值value是否在指定的slice中存在
func InUint32Slices(value uint32, slices []uint32) bool

// InUint64Slices 判断指定值value是否在指定的slice中存在
func InUint64Slices(value uint64, slices []uint64) bool

// MergeSlices 将多个slice合并成一个slice
func MergeSlices(slices ...[]interface{}) []interface{}

// MergeStrSlices 将多个slice合并成一个slice
func MergeStrSlices(slices ...[]string) []string

// MergeIntSlices 将多个slice合并成一个slice
func MergeIntSlices(slices ...[]int) []int

// MergeInt8Slices 将多个slice合并成一个slice
func MergeInt8Slices(slices ...[]int8) []int8

// MergeInt16Slices 将多个slice合并成一个slice
func MergeInt16Slices(slices ...[]int16) []int16

// MergeInt32Slices 将多个slice合并成一个slice
func MergeInt32Slices(slices ...[]int32) []int32 

// MergeInt64Slices 将多个slice合并成一个slice
func MergeInt64Slices(slices ...[]int64) []int64

// MergeUintSlices 将多个slice合并成一个slice
func MergeUintSlices(slices ...[]uint) []uint

// MergeUint8Slices 将多个slice合并成一个slice
func MergeUint8Slices(slices ...[]uint8) []uint8

// MergeUint16Slices 将多个slice合并成一个slice
func MergeUint16Slices(slices ...[]uint16) []uint16

// MergeUint32Slices 将多个slice合并成一个slice
func MergeUint32Slices(slices ...[]uint32) []uint32

// MergeUint64Slices 将多个slice合并成一个slice
func MergeUint64Slices(slices ...[]uint64) []uint64

// SumIntSlice 对slice中的元素求和
func SumIntSlice(slice []int) int

// SumInt8Slice 对slice中的元素求和
func SumInt8Slice(slice []int8) int8

// SumInt16Slice 对slice中的元素求和
func SumInt16Slice(slice []int16) int16

// SumInt32Slice 对slice中的元素求和
func SumInt32Slice(slice []int32) int32

// SumInt64Slice 对slice中的元素求和
func SumInt64Slice(slice []int64) int64

// SumFloat32Slice 对slice中的元素求和
func SumFloat32Slice(slice []float32) float32

// SumFloat64Slice 对slice中的元素求和
func SumFloat64Slice(slice []float64) float64
```

### randUtil 随机数

```go
// RandomInt 生成指定范围内的随机整数
func RandomInt(min, max int) int

// RandomFloat 生成指定范围内的随机浮点数
func RandomFloat(min, max float64) float64

// RandomString 生成指定长度的随机字符串
func RandomString(length int) string

// RandomBytes 生成指定长度的随机字节数组
func RandomBytes(length int) []byte
```

### mapUtil map类型处理

```go
// MapValueExists 判断map中的value是否存在
func MapValueExists(m map[string]interface{}, value interface{}) bool

// MapKeyExists 判断map中的key是否存在
func MapKeyExists(m map[string]interface{}, key string) bool
```

### intUtil 数值型处理

```go
// StrToInt 将string类型转换为int类型
func StrToInt(s string) (int, error)

// StrToInt8 将string类型转换为int8类型
func StrToInt8(s string) (int8, error)

// StrToInt16 将string类型转换为int16类型
func StrToInt16(s string) (int16, error)

// StrToInt32 将string类型转换为int32类型
func StrToInt32(s string) (int32, error)

// StrToInt64 将string类型转换为int64类型
func StrToInt64(s string) (int64, error)

// IntToString 将int类型转换为string类型
func IntToString(n int) string

// Int8ToString 将int8类型转换为string类型
func Int8ToString(n int8) string

// Int16ToString 将int16类型转换为string类型
func Int16ToString(n int16) string

// Int32ToString 将int32类型转换为string类型
func Int32ToString(n int32) string

// Int64ToString 将int64类型转换为string类型
func Int64ToString(n int64) string
```

### floatUtil 浮点型处理

```go
// Float32ToStr float32转字符串
func Float32ToStr(f float32) string

// Float64ToStr float64转字符串
func Float64ToStr(input float64) string

// Float32ToFloat64 float32转float64
func Float32ToFloat64(f float32) float64

// Float64ToFloat32 float64转float32
func Float64ToFloat32(f float64) float32

// Float64Round 四舍五入
func Float64Round(input float64, precision int) float64

// Float64Ceil 向上取整
func Float64Ceil(input float64) float64
```

### emoji表情包

```go
// UnicodeEmojiDecode Emoji表情解码
func UnicodeEmojiDecode(s string) string

// UnicodeEmojiCode Emoji表情转换
func UnicodeEmojiCode(s string) string
```

### cryptoUtil 加密与解密

```go
// HashSHA256 hash加密
func HashSHA256(str string) string

// Md5 MD5加密
func Md5(string string) string
```

### byteUtil 字节数组

```go
// BytesToStr 字节数组转字符串
func BytesToStr(data []byte) string
```

### anyUtil 任意类型转换

```go
// AnyToFloat32 将给定的值转换为float32
func AnyToFloat32(value interface{}) (float32, error) 

// AnyToFloat64 将给定的值转换为float64
func AnyToFloat64(v interface{}) (float64, error)

// AnyToInt 将给定的值转换为 int
func AnyToInt(input interface{}) (int, error)

// AnyToInt8 将给定的值转换为 int8
func AnyToInt8(input interface{}) (int8, error)

// AnyToInt16 将给定的值转换为 int16
func AnyToInt16(input interface{}) (int16, error)

// AnyToInt32 将给定的值转换为 int32
func AnyToInt32(input interface{}) (int32, error)

// AnyToInt64 将给定的值转换为 int64
func AnyToInt64(input interface{}) (int64, error)

// AnyToStr 任意类型数据转string
func AnyToStr(input interface{}) string 

// AnyToUint 将给定的值转换为 uint
func AnyToUint(input interface{}) (uint, error)

// AnyToUint8 将给定的值转换为 uint8
func AnyToUint8(input interface{}) (uint8, error)

// AnyToUint16 将给定的值转换为 uint16
func AnyToUint16(input interface{}) (uint16, error)

// AnyToUint32 将给定的值转换为 uint32
func AnyToUint32(input interface{}) (uint32, error)

// AnyToUint64 将给定的值转换为 uint64
func AnyToUint64(input interface{}) (uint64, error)
```


## 赞助与支持

`GoEasyUtils ` Thank JetBrains for their support

<a href="https://www.jetbrains.com"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" height="100" alt="JetBrains"/></a>