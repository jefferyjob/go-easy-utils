---
name: "\U0001F41B Bug report"
about: Report a reproducible bug or regression.
labels: Bug
---
Hello,

I encountered an issue with the following code:
```go
var slice = []string{"this", "is", "go", "easy", "utils"}
chunkSlice := sliceUtil.ChunkStr(slice, 2)
```

golang version: **such as 1.16**

carbon version: **such as 1.0.1**

I expected to get:

```
[]string{"this", "is"}
[]string{"go", "easy"}
[]string{"utils"}
```

But I actually get:

```
[]string{"this", "is"}
[]string{"go", "easy", "utils"}
```

Thanks!