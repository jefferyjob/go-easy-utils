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
go get -u github.com/jefferyjob/go-easy-utils/v3
```

**Use Demo**
```go
package main

import (
	"fmt"
	"github.com/jefferyjob/go-easy-utils/v3/slicex"
)

func main() {
	var slice = []string{"this", "is", "go", "easy", "utils"}
	chunkSlice := slicex.ChunkSlice(slice, 2)
	fmt.Printf("%v", chunkSlice)
}
```

## Function list

| Package name | Function Outline                                                                          | Document             |
|--------------| ----------------------------------------------------------------------------------------- |----------------------|
| anyx         | Convert any type of data to the specified type                                            | [README](anyx)    |
| bytex        | Conversion of byte array                                                                  | [README](bytex)   |
| cryptox      | Various encryption processing                                                             | [README](cryptox) |
| emojix       | Decoding and encoding of emoji expression                                                 | [README](emojix)  |
| floatx       | Floating-point data processing                                                            | [README](floatx)  |
| intx         | Numerical data processing                                                                 | [README](intUtil)    |
| jsonx        | Json data conversion, support weak type conversion                                        | [README](jsonx)   |
| mapx         | Map type data processing                                                                  | [README](mapx)    |
| mathx        | The Math function can handle values within the range of integers and floats.              | [README](mathx)   |
| randx        | Random number generation, including: number, string, byte array                           | [README](randUtil)   |
| slicex       | Slice processing (grouping, summation, transformation, merging, etc.)                     | [README](slicex)  |
| strx         | String conversion processing                                                              | [README](strx)    |
| validx       | Common data verification, such as: Chinese, English, name, ID number, phone number, email | [README](validx)  |


## License
This library is licensed under the Apache-2.0. See the LICENSE file for details.