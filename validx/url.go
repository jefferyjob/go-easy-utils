package validx

import (
	"regexp"
)

// IsURL 是否为URL地址
func IsURL(url string) bool {
	match, _ := regexp.MatchString(`^(http|https)://[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?$`, url)
	return match
}
