# mathUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/mathUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/mathUtil"
)
```

## Functions

```go
// Abs 返回一个数的绝对值
func Abs[T Numeric](num T) T

// Ceil 对float数据向上取整
func Ceil[T float32 | float64](num T) int

// Floor 对float数据向下取整
func Floor[T float32 | float64](num T) int

// Max 返回slice中最大值
func Max[T Numeric](slice []T) T

// Min 返回slice中最小值
func Min[T Numeric](slice []T) T

// Round 对float数据四舍五入
func Round[T float32 | float64](num T) int
```