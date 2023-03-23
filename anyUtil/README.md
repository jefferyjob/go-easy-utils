# anyUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/anyUtil
```

## ErrorCode

```go
// ErrSyntax indicates that a value does not have the right syntax for the target type
// 指示值不具有目标类型的正确语法
var ErrSyntax = strconv.ErrSyntax

// ErrType indicates that a value does not have the right syntax for the target type
// 指示值不具有目标类型的正确语法
var ErrType = errors.New("unsupported type")

// ErrValOut Indicates that the range corresponding to the type is exceeded
// 表示超出了类型对应的范围
var ErrValOut = errors.New("value out of range")

// ErrUnsignedInt The identity type is negative, so it cannot be converted to an unsigned integer
// 标识类型为负数，所以无法转为无符号的整型
var ErrUnsignedInt = errors.New("cannot convert negative value to unsigned integer")
```

## Functions

```go
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