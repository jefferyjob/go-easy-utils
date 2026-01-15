# strx

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/v3/strx
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/v3/strx"
)
```

## Functions

```go
// ToInt string转int
func ToInt(v string) int 

// ToInt8 string转int8
func ToInt8(v string) int8 

// ToInt16 string转int16
func ToInt16(v string) int16 

// ToInt32 string转int32
func ToInt32(v string) int32 

// ToInt64 string转int64
func ToInt64(v string) int64

// ToUint string转uint
func ToUint(v string) uint

// ToUint8 string转uint8
func ToUint8(v string) uint8

// ToUint16 string转uint16
func ToUint16(v string) uint16

// ToUint32 string转uint32
func ToUint32(v string) uint32

// ToUint64 string转uint64
func ToUint64(v string) uint64

// ToBytes 字符串转字节数组
func ToBytes(v string) []byte

// Cut 删除 s 中出现的 sub 字符串
func Cut(s, sub string, n ...int) string

// 判断权限路径模式是否可以匹配实际请求路径
func PathMatch(pattern, path string) bool
```