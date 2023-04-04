# emoji表情包

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/emojiUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/emojiUtil"
)
```

## Functions

```go
// DecodeEmojiUnicode Emoji表情解码
func DecodeEmojiUnicode(unicode string) string

// EncodeEmojiUnicode Emoji表情编码
func EncodeEmojiUnicode(emoji string) string
```

### emoji表情unicode编码表

- https://unicode.org/charts/
- https://home.unicode.org/emoji/emoji-frequency/