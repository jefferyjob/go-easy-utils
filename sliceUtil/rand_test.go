package sliceUtil

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	arr := []interface{}{"apple", "banana", "orange", "grape", "peach"}
	fmt.Println(Rand(arr, 3))
}
