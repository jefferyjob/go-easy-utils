# go-easy-utils

[![Go](https://img.shields.io/badge/Go-%3E%3D1.18-green)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/jefferyjob/go-easy-utils.svg)](https://github.com/jefferyjob/go-easy-utils/releases)
[![Action](https://github.com/jefferyjob/go-easy-utils/workflows/Go/badge.svg?branch=main)](https://github.com/jefferyjob/go-easy-utils/actions)
[![Report](https://goreportcard.com/badge/github.com/jefferyjob/go-easy-utils)](https://goreportcard.com/report/github.com/jefferyjob/go-easy-utils)
[![Coverage](https://codecov.io/gh/jefferyjob/go-easy-utils/branch/main/graph/badge.svg)](https://codecov.io/gh/jefferyjob/go-easy-utils)
[![Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/jefferyjob/go-easy-utils/v2)
[![License](https://img.shields.io/github/license/jefferyjob/go-easy-utils)](https://github.com/jefferyjob/go-easy-utils/blob/main/LICENSE)

[English](README.md) | 简体中文

## 介绍
这是一个基于 Go 语言开发的通用数据类型处理工具类，帮助开发者在业务代码实现中处理常见的数据类型和数据操作。可以让您专注于您的业务代码的实现，而免去处理基本数据类型转换和验证的功能。该工具库无侵入式的设计可以让您的业务代码更容易阅读和优雅。

## 快速开始
**安装**

使用 `Go1.18` 及以上版本的用户，建议安装 `v2.x.x`。 因为 `v2.x.x` 应用 `Go1.18` 的泛型重写了大部分函数。
```bash
go get -u github.com/jefferyjob/go-easy-utils/v2
```

使用 `Go1.18` 以下版本的用户，请使用 [v1.x](https://github.com/jefferyjob/go-easy-utils/tree/v1) 版本

**使用Demo**
```go
package main

import (
	"fmt"
	"github.com/jefferyjob/go-easy-utils/v2/sliceUtil"
)

func main() {
	var slice = []string{"this", "is", "go", "easy", "utils"}
	chunkSlice := sliceUtil.ChunkSlice(slice, 2)
	fmt.Printf("%v", chunkSlice)
}
```

## 功能列表

### jsonUtil Json处理工具

```go
// JsonToStruct 将 JSON 字符串解析为指定的结构体指针
func JsonToStruct(jsonData string, result any) error
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
func StrToInt(v string) int

// StrToInt8 string转int8
func StrToInt8(v string) int8

// StrToInt16 string转int16
func StrToInt16(v string) int16

// StrToInt32 string转int32
func StrToInt32(v string) int32

// StrToInt64 string转int64
func StrToInt64(v string) int64

// StrToUint string转uint
func StrToUint(v string) uint

// StrToUint8 string转uint8
func StrToUint8(v string) uint8

// StrToUint16 string转uint16
func StrToUint16(v string) uint16

// StrToUint32 string转uint32
func StrToUint32(v string) uint32

// StrToUint64 string转uint64
func StrToUint64(v string) uint64

// StrToBytes 字符串转字节数组
func StrToBytes(v string) []byte 
```

### sliceUtil 切片处理工具

```go
// Chunk 把slice分割为新的数组块
func ChunkSlice(slice []T, size int) [][]T

// Column 获取slice中某个单一列的值
func ColumnSlice(slice []T, column string) []any

// In 判断value是否在slice中
func InSlice(value T, slices []T) bool

// Is 判断指定值i是否是slice类型
func IsSlice(slice any) bool

// Merge 将多个slice合并成一个slice
func MergeSlice(slices ...[]T) []T

// Sum 对slice中的元素求和
func SumSlice(slice []T) T

// Unique 移除slice中重复的值
func UniqueSlice(slice []T) []T
```

### mapUtil map类型处理

```go
// MapKeyExists 判断map中的key是否存在
func MapKeyExists((m map[T]T2, key T)) bool

// MapValueExists 判断map中的value是否存在
func MapValueExists(m map[T2]T, value T) bool
```

### mathUtil

```go
// Abs 返回一个数的绝对值
func Abs(num T) T

// Ceil 对float数据向上取整
func Ceil(num T) int

// Floor 对float数据向下取整
func Floor(num T) int

// Max 返回slice中最大值
func Max(slice []T) T

// Min 返回slice中最小值
func Min(slice []T) T

// Round 对float数据四舍五入
func Round(num T) int
```

### floatUtil 浮点型处理

```go
// Float32ToStr float32转字符串
func Float32ToStr(f float32) string

// Float64ToStr float64转字符串
func Float64ToStr(f float64) string

// Float32ToFloat64 float32转float64
func Float32ToFloat64(f float32) float64

// Float64ToFloat32 float64转float32
func Float64ToFloat32(f float64) float32
```

### emoji表情包

```go
// DecodeEmojiUnicode Emoji表情解码
func DecodeEmojiUnicode(unicode string) string

// EncodeEmojiUnicode Emoji表情编码
func EncodeEmojiUnicode(emoji string) string
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
func AnyToFloat32(i any) (float32, error)

// AnyToFloat64 将给定的值转换为float64
func AnyToFloat64(i any) (float64, error)

// AnyToInt 将给定的值转换为 int
func AnyToInt(i any) (int, error)

// AnyToInt8 将给定的值转换为 int8
func AnyToInt8(i any) (int8, error)

// AnyToInt16 将给定的值转换为 int16
func AnyToInt16(i any) (int16, error)

// AnyToInt32 将给定的值转换为 int32
func AnyToInt32(i any) (int32, error)

// AnyToInt64 将给定的值转换为 int64
func AnyToInt64(i any) (int64, error)

// AnyToStr 任意类型数据转string
func AnyToStr(i any) string

// AnyToUint 将给定的值转换为 uint
func AnyToUint(i any) (uint, error)

// AnyToUint8 将给定的值转换为 uint8
func AnyToUint8(i any) (uint8, error)

// AnyToUint16 将给定的值转换为 uint16
func AnyToUint16(i any) (uint16, error)

// AnyToUint32 将给定的值转换为 uint32
func AnyToUint32(i any) (uint32, error)

// AnyToUint64 将给定的值转换为 uint64
func AnyToUint64(i any) (uint64, error)

// AnyToBool 将给定的值转换为bool
func AnyToBool(i any) bool 
```



