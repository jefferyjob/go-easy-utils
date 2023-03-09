package debugUtil

import (
	"fmt"
	"os"
)

// DD Debug打断调试
func DD(value ...interface{}) {
	fmt.Println(value)
	os.Exit(1)
}
