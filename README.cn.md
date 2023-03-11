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

## 快速上手

**安装**

Go version >= 1.16 (recommend)

```bash
go get -u github.com/jefferyjob/go-easy-utils

import (
	"github.com/jefferyjob/go-easy-utils"
)
```

**使用**

导入包
```
import "github.com/jefferyjob/go-easy-utils"
```

使用包
```

```

## 功能列表

### anyUtil 任意类型转换

```
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

### byteUtil 字节数组

```
// BytesToStr 字节数组转字符串
func BytesToStr(data []byte) string
```

### cryptoUtil 加密与解密

```
// HashSHA256 hash加密
func HashSHA256(str string) string

// Md5 MD5加密
func Md5(string string) string
```

## 赞助与支持

`GoEasyUtils ` Thank JetBrains for their support

<a href="https://www.jetbrains.com"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" height="100" alt="JetBrains"/></a>