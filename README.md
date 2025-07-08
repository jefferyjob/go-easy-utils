# go-easy-utils

[![Go](https://img.shields.io/badge/Go->=1.24-green)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/jefferyjob/go-easy-utils.svg)](https://github.com/jefferyjob/go-easy-utils/releases)
[![Action](https://github.com/jefferyjob/go-easy-utils/workflows/Go/badge.svg?branch=main)](https://github.com/jefferyjob/go-easy-utils/actions)
[![Report](https://goreportcard.com/badge/github.com/jefferyjob/go-easy-utils)](https://goreportcard.com/report/github.com/jefferyjob/go-easy-utils)
[![Coverage](https://codecov.io/gh/jefferyjob/go-easy-utils/branch/main/graph/badge.svg)](https://codecov.io/gh/jefferyjob/go-easy-utils)
[![Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/jefferyjob/go-easy-utils/v2)
[![License](https://img.shields.io/github/license/jefferyjob/go-easy-utils)](https://github.com/jefferyjob/go-easy-utils/blob/main/LICENSE)

English | [简体中文](README.cn.md)

## Introduction
This is a general data type processing tool class based on Go language, which helps developers process common data types and data operations in business code implementation. It allows you to focus on the implementation of your business code without processing the basic data type conversion and validation functions. The non-intrusive design of the tool library can make your business code easier to read and elegant.

## Quick Start
**Install**
```bash
go get -u github.com/jefferyjob/go-easy-utils/v2
```

Users who use versions below `Go1.18`, please use [v1.x](https://github.com/jefferyjob/go-easy-utils/tree/v1) version

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

`go-easy-utils` Thank JetBrains for their support

<a href="https://www.jetbrains.com"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" height="100" alt="JetBrains"/></a>

## License
This library is licensed under the Apache-2.0. See the LICENSE file for details.