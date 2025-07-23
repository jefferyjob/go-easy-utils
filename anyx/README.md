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

```
// ToFloat32 将给定的值转换为float32
ToFloat32(i any) (float32, error) 

// ToFloat64 将给定的值转换为float64
ToFloat64(i any) (float64, error)

// ToInt 将给定的值转换为 int
ToInt(i any) (int, error)

// ToInt8 将给定的值转换为 int8
ToInt8(i any) (int8, error)

// ToInt16 将给定的值转换为 int16
ToInt16(i any) (int16, error)

// ToInt32 将给定的值转换为 int32
ToInt32(i any) (int32, error)

// ToInt64 将给定的值转换为 int64
ToInt64(i any) (int64, error)

// ToStr 任意类型数据转string
ToStr(i any) string 

// ToUint 将给定的值转换为 uint
ToUint(i any) (uint, error)

// ToUint8 将给定的值转换为 uint8
ToUint8(i any) (uint8, error)

// ToUint16 将给定的值转换为 uint16
ToUint16(i any) (uint16, error)

// ToUint32 将给定的值转换为 uint32
ToUint32(i any) (uint32, error)

// ToUint64 将给定的值转换为 uint64
ToUint64(i any) (uint64, error)

// ToBool 将给定的值转换为bool
ToBool(i any) bool

// Ternary 三元运算符
Ternary(expr bool, a, b T) T 
```