package sliceUtil

import (
	"math/rand"
	"time"
)

func Rand(slice []interface{}, num int) []int {
	rand.Seed(time.Now().UnixNano())
	keys := make([]int, num)
	for i := 0; i < num; i++ {
		keys[i] = rand.Intn(len(slice))
	}
	return keys
}
