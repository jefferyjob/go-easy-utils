package sliceUtil

import (
	"fmt"
	"testing"
)

func TestMergeStr(t *testing.T) {
	var s1 = []string{"a", "b"}
	var s2 = []string{"a"}
	var s3 = []string{"a"}

	res := MergeStr(s1, s2, s3)

	fmt.Printf("%+v", res)
}
