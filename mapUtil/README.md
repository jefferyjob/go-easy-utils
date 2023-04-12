# mapUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/mapUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/mapUtil"
)
```

## Functions

```go
// MapValueExists 判断map中的value是否存在
func MapValueExists(m map[string]T, value T) bool

// MapKeyExists 判断map中的key是否存在
func MapKeyExists(m map[string]any, key string) bool
```