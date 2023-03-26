# sliceUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/sliceUtil
```

## Functions

```go
// Chunk 将一个切片按指定大小分成多个切片，并返回一个包含这些切片的二维切片。
func Chunk(slice []interface{}, size int) [][]interface{} 

// ChunkStr 将一个 string 类型的切片切割成指定大小的多个子切片
func ChunkStr(slice []string, size int) [][]string

// ChunkInt 将一个 int 类型的切片切割成指定大小的多个子切片
func ChunkInt(slice []int, size int) [][]int 

// ChunkInt8 将一个 int8 类型的切片切割成指定大小的多个子切片
func ChunkInt8(slice []int8, size int) [][]int8

// ChunkInt32 将一个 int32 类型的切片切割成指定大小的多个子切片
func ChunkInt32(slice []int32, size int) [][]int32

// ChunkInt64 将一个 int64 类型的切片切割成指定大小的多个子切片
func ChunkInt64(slice []int64, size int) [][]int64

// ChunkUint 将一个 uint 类型的切片切割成指定大小的多个子切片
func ChunkUint(slice []uint, size int) [][]uint

// ChunkUint8 将一个 uint8 类型的切片切割成指定大小的多个子切片
func ChunkUint8(slice []uint8, size int) [][]uint8

// ChunkUint16 将一个 uint16 类型的切片切割成指定大小的多个子切片
func ChunkUint16(slice []uint16, size int) [][]uint16

// ChunkUint32 将一个 uint32 类型的切片切分成固定大小的子切片
func ChunkUint32(slice []uint32, size int) [][]uint32

// ChunkUint64 将一个 uint64 类型的切片切分成固定大小的子切片
func ChunkUint64(slice []uint64, size int) [][]uint64

// InSlices 判断指定值value是否在指定的slice中存在
func InSlices(value interface{}, slices []interface{}) bool

// InStrSlices 判断指定值value是否在指定的slice中存在
func InStrSlices(value string, slices []string) bool 

// InIntSlices 判断指定值value是否在指定的slice中存在
func InIntSlices(value int, slices []int) bool

// InInt8Slices 判断指定值value是否在指定的slice中存在
func InInt8Slices(value int8, slices []int8) bool

// InInt16Slices 判断指定值value是否在指定的slice中存在
func InInt16Slices(value int16, slices []int16) bool 

// InInt32Slices 判断指定值value是否在指定的slice中存在
func InInt32Slices(value int32, slices []int32) bool

// InInt64Slices 判断指定值value是否在指定的slice中存在
func InInt64Slices(value int64, slices []int64) bool 

// InUintSlices 判断指定值value是否在指定的slice中存在
func InUintSlices(value uint, slices []uint) bool

// InUint8Slices 判断指定值value是否在指定的slice中存在
func InUint8Slices(value uint8, slices []uint8) bool

// InUint16Slices 判断指定值value是否在指定的slice中存在
func InUint16Slices(value uint16, slices []uint16) bool

// InUint32Slices 判断指定值value是否在指定的slice中存在
func InUint32Slices(value uint32, slices []uint32) bool

// InUint64Slices 判断指定值value是否在指定的slice中存在
func InUint64Slices(value uint64, slices []uint64) bool

// MergeSlices 将多个slice合并成一个slice
func MergeSlices(slices ...[]interface{}) []interface{}

// MergeStrSlices 将多个slice合并成一个slice
func MergeStrSlices(slices ...[]string) []string

// MergeIntSlices 将多个slice合并成一个slice
func MergeIntSlices(slices ...[]int) []int

// MergeInt8Slices 将多个slice合并成一个slice
func MergeInt8Slices(slices ...[]int8) []int8

// MergeInt16Slices 将多个slice合并成一个slice
func MergeInt16Slices(slices ...[]int16) []int16

// MergeInt32Slices 将多个slice合并成一个slice
func MergeInt32Slices(slices ...[]int32) []int32 

// MergeInt64Slices 将多个slice合并成一个slice
func MergeInt64Slices(slices ...[]int64) []int64

// MergeUintSlices 将多个slice合并成一个slice
func MergeUintSlices(slices ...[]uint) []uint

// MergeUint8Slices 将多个slice合并成一个slice
func MergeUint8Slices(slices ...[]uint8) []uint8

// MergeUint16Slices 将多个slice合并成一个slice
func MergeUint16Slices(slices ...[]uint16) []uint16

// MergeUint32Slices 将多个slice合并成一个slice
func MergeUint32Slices(slices ...[]uint32) []uint32

// MergeUint64Slices 将多个slice合并成一个slice
func MergeUint64Slices(slices ...[]uint64) []uint64

// SumIntSlice 对slice中的元素求和
func SumIntSlice(slice []int) int

// SumInt8Slice 对slice中的元素求和
func SumInt8Slice(slice []int8) int8

// SumInt16Slice 对slice中的元素求和
func SumInt16Slice(slice []int16) int16

// SumInt32Slice 对slice中的元素求和
func SumInt32Slice(slice []int32) int32

// SumInt64Slice 对slice中的元素求和
func SumInt64Slice(slice []int64) int64

// SumFloat32Slice 对slice中的元素求和
func SumFloat32Slice(slice []float32) float32

// SumFloat64Slice 对slice中的元素求和
func SumFloat64Slice(slice []float64) float64
```