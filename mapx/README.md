# mapx

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/v2/mapUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/v2/mapUtil"
)
```

## Functions

```go
// KeyExists 判断map中的key是否存在
func KeyExists((m map[T]T2, key T)) bool

// ValueExists 判断map中的value是否存在
func ValueExists(m map[T2]T, value T) bool

// Keys 提取的所有键
func Keys(m map[K]V) []K

// Values 提取的所有值
func Values(m map[K]V) []V
```