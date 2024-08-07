# sliceUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/v2/sliceUtil
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

// SliceToMap 切片转字典
func SliceToMap[T any, K comparable](items []T, keyFunc func(T) K) map[K]T

// ExtractKeys 字段提取
func ExtractKeys[T any, K comparable](items []T, keyFunc func(T) K) []K

// Diff 计算差集
func Diff[T comparable](s []T, slices ...[]T) []T

// SymmetricDiff 计算对称差集
func SymmetricDiff[T comparable](slices ...[]T) []T

// Intersect 计算交集
func Intersect[T comparable](slices ...[]T) []T
```