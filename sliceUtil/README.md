# sliceUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/sliceUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/sliceUtil"
)
```

## Functions

```go
// Chunk 把slice分割为新的数组块
func ChunkSlice[T any](slice []T, size int) [][]T

// Column 获取slice中某个单一列的值
func ColumnSlice[T any](slice []T, column string) []any

// In 判断value是否在slice中
func InSlice[T comparable](value T, slices []T) bool

// Is 判断指定值i是否是slice类型
func IsSlice(slice any) bool

// Merge 将多个slice合并成一个slice
func MergeSlice[T any](slices ...[]T) []T

// Sum 对slice中的元素求和
func SumSlice[T Numeric](slice []T) T

// Unique 移除slice中重复的值
func UniqueSlice[T comparable](slice []T) []T
```