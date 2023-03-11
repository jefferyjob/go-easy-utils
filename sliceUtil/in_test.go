package sliceUtil

import (
	"fmt"
	"testing"
)

func TestInStr(t *testing.T) {
	var slices1 = []string{"a", "b", "c"}
	var val1 = "a"
	res1 := InStrSlices(val1, slices1)
	res2 := InStrSlices("abc", slices1)
	fmt.Println("结果判定", res1, res2)
}
