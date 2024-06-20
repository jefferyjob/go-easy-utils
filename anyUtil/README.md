# anyUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/v2/anyUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/v2/anyUtil"
)
```

## Functions

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