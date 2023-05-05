# sliceUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/sliceUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/v2/sliceUtil"
)
```

## Functions

```go
// Chunk 把slice分割为新的数组块
func ChunkSlice(slice []T, size int) [][]T

// Column 获取slice中某个单一列的值
func ColumnSlice(slice []T, column string) []any

// In 判断value是否在slice中
func InSlice(value T, slices []T) bool

// Is 判断指定值i是否是slice类型
func IsSlice(slice any) bool

// Merge 将多个slice合并成一个slice
func MergeSlice(slices ...[]T) []T

// Sum 对slice中的元素求和
func SumSlice(slice []T) T

// Unique 移除slice中重复的值
func UniqueSlice(slice []T) []T
```