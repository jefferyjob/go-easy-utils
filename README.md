# go-easy-utils

[![Go](https://img.shields.io/badge/Go-%3E%3D1.18-green)](https://go.dev)
[![Release](https://img.shields.io/github/release/jefferyjob/go-easy-utils.svg)](https://github.com/jefferyjob/go-easy-utils/releases)
[![Go Action](https://github.com/jefferyjob/go-easy-utils/workflows/Go/badge.svg?branch=master)](https://github.com/jefferyjob/go-easy-utils/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/jefferyjob/go-easy-utils)](https://goreportcard.com/report/github.com/jefferyjob/go-easy-utils)
[![Go Coverage](https://codecov.io/gh/jefferyjob/go-easy-utils/branch/master/graph/badge.svg)](https://codecov.io/gh/jefferyjob/go-easy-utils)
[![GoEasyUtils Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/jefferyjob/go-easy-utils/v2)
![License](https://img.shields.io/github/license/jefferyjob/go-easy-utils)

English | [简体中文](README.cn.md)

## Introduction
This is a general data type processing tool class based on Go language, which helps developers process common data types and data operations in business code implementation. It allows you to focus on the implementation of your business code without processing the basic data type conversion and validation functions. The non-intrusive design of the tool library can make your business code easier to read and elegant.

## Quick Start
**Install**

Use users with `Go1.18` and above, it is recommended to install `v2.x.x`. Because `v2.x.x` app rewritten most functions of `Go1.18`
```bash
go get -u github.com/jefferyjob/go-easy-utils/v2
```

Users who use `Go1.18` below must install `v1.x.x`. The latest `v1` version is `v1.1.0`
```bash
go get github.com/jefferyjob/go-easy-utils@v1.1.0
```

**Use Demo**
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

## Function list

| Package name | Function Outline                                                                          | Document             |
|--------------| ----------------------------------------------------------------------------------------- |----------------------|
| anyUtil      | Convert any type of data to the specified type                                            | [README](anyUtil)    |
| byteUtil     | Conversion of byte array                                                                  | [README](byteUtil)   |
| cryptoUtil   | Various encryption processing                                                             | [README](cryptoUtil) |
| emojiUtil    | Decoding and encoding of emoji expression                                                 | [README](emojiUtil)  |
| floatUtil    | Floating-point data processing                                                            | [README](floatUtil)  |
| intUtil      | Numerical data processing                                                                 | [README](intUtil)    |
| jsonUtil     | Json data conversion, support weak type conversion                                        | [README](jsonUtil)   |
| mapUtil      | Map type data processing                                                                  | [README](mapUtil)    |
| mathUtil     | The Math function can handle values within the range of integers and floats.              | [README](mathUtil)   |
| randUtil     | Random number generation, including: number, string, byte array                           | [README](randUtil)   |
| sliceUtil    | Slice processing (grouping, summation, transformation, merging, etc.)                     | [README](sliceUtil)  |
| strUtil      | String conversion processing                                                              | [README](strUtil)    |
| validUtil    | Common data verification, such as: Chinese, English, name, ID number, phone number, email | [README](validUtil)  |


## Sponsorship and support

`GoEasyUtils` Thank JetBrains for their support

<a href="https://www.jetbrains.com"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" height="100" alt="JetBrains"/></a>