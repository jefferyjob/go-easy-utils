# go-easy-utils

[![Go](https://img.shields.io/badge/Go->=1.24-green)](https://go.dev)
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
```bash
go get -u github.com/jefferyjob/go-easy-utils/v2
```

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

| 包名         | 函数概要                                                                                  | 文档                 |
|--------------| ----------------------------------------------------------------------------------------- |----------------------|
| anyUtil      | 将任意类型的数据转换为指定类型                                                            | [README](anyUtil)    |
| byteUtil     | 字节数组转换                                                                             | [README](byteUtil)   |
| cryptoUtil   | 各种加密处理                                                                             | [README](cryptoUtil) |
| emojiUtil    | 表情符号的解码和编码                                                                     | [README](emojiUtil)  |
| floatUtil    | 浮点数数据处理                                                                           | [README](floatUtil)  |
| intUtil      | 数值数据处理                                                                           | [README](intUtil)    |
| jsonUtil     | JSON 数据转换，支持弱类型转换                                                             | [README](jsonUtil)   |
| mapUtil      | Map 类型数据处理                                                                        | [README](mapUtil)    |
| mathUtil     | 数学函数可以处理整数和浮点数范围内的值                                                     | [README](mathUtil)   |
| randUtil     | 随机数生成，包括：数字、字符串、字节数组                                                   | [README](randUtil)   |
| sliceUtil    | 切片处理（分组、求和、转换、合并等）                                                       | [README](sliceUtil)  |
| strUtil      | 字符串转换处理                                                                           | [README](strUtil)    |
| validUtil    | 常见数据验证，如：中文、英文、姓名、身份证号、电话号码、电子邮件                          | [README](validUtil)  |

## 许可证
本库采用 Apache-2.0 进行授权。有关详细信息，请参阅 LICENSE 文件。
