# mapUtil

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
// MapKeyExists 判断map中的key是否存在
func MapKeyExists((m map[T]T2, key T)) bool

// MapValueExists 判断map中的value是否存在
func MapValueExists(m map[T2]T, value T) bool

// ExtractKeys 提取Map的所有键
func ExtractKeys[K comparable, V any](m map[K]V) []K

// ExtractValues 提取Map的所有值
func ExtractValues[K comparable, V any](m map[K]V) []V
```