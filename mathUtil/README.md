# mathUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/mathUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/v2/mathUtil"
)
```

## Functions

```go
// Abs 返回一个数的绝对值
func Abs(num T) T

// Ceil 对float数据向上取整
func Ceil(num T) int

// Floor 对float数据向下取整
func Floor(num T) int

// Max 返回slice中最大值
func Max(slice []T) T

// Min 返回slice中最小值
func Min(slice []T) T

// Round 对float数据四舍五入
func Round(num T) int
```