# anyx

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/v3/anyx
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/v3/anyx"
)
```

## Functions

```go
// ToFloat32 将给定的值转换为float32
func ToFloat32(i any) (float32, error) 

// ToFloat64 将给定的值转换为float64
func ToFloat64(i any) (float64, error)

// ToInt 将给定的值转换为 int
func ToInt(i any) (int, error)

// ToInt8 将给定的值转换为 int8
func ToInt8(i any) (int8, error)

// ToInt16 将给定的值转换为 int16
func ToInt16(i any) (int16, error)

// ToInt32 将给定的值转换为 int32
func ToInt32(i any) (int32, error)

// ToInt64 将给定的值转换为 int64
func ToInt64(i any) (int64, error)

// ToStr 任意类型数据转string
func ToStr(i any) string 

// ToUint 将给定的值转换为 uint
func ToUint(i any) (uint, error)

// ToUint8 将给定的值转换为 uint8
func ToUint8(i any) (uint8, error)

// ToUint16 将给定的值转换为 uint16
func ToUint16(i any) (uint16, error)

// ToUint32 将给定的值转换为 uint32
func ToUint32(i any) (uint32, error)

// ToUint64 将给定的值转换为 uint64
func ToUint64(i any) (uint64, error)

// ToBool 将给定的值转换为bool
func ToBool(i any) bool

// Ternary 三元运算符
func Ternary[T any](expr bool, a, b T) T 
```