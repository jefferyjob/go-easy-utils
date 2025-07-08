# sliceUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/v3/slicex
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/v3/slicex"
)
```

## Functions

```go
// Chunk 把slice分割为新的数组块
func Chunk(slice []T, size int) [][]T

// Column 获取slice中某个单一列的值
func Column(slice []T, column string) []any

// In 判断value是否在slice中
func In(value T, slices []T) bool

// Is 判断指定值i是否是slice类型
func Is(slice any) bool

// Merge 将多个slice合并成一个slice
func Merge(slices ...[]T) []T

// Sum 对slice中的元素求和
func Sum(slice []T) T

// Unique 移除slice中重复的值
func Unique(slice []T) []T

// ToMap 切片转map
func ToMap[T any, K comparable](items []T, keyFunc func(T) K) map[K]T

// Pluck 字段提取
func Pluck[T any, K comparable](items []T, keyFunc func(T) K) []K

// Diff 计算差集
func Diff[T comparable](s []T, slices ...[]T) []T

// SymmetricDiff 计算对称差集
func SymmetricDiff[T comparable](slices ...[]T) []T

// Intersect 计算交集
func Intersect[T comparable](slices ...[]T) []T
```