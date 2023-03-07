// 提供了一个名为Rand的函数，用于从给定的切片中返回指定数量的随机键。
//
// 函数签名如下：
//
// 	func Rand(slice []interface{}, num int) []int
//
// 参数slice是要从中获取随机键的切片，参数num是要返回多少个随机键。
//
// 返回值是一个包含多个随机键的切片。如果num大于slice的长度，函数会返回slice的所有键。
//
// 示例：
//
// 	arr := []interface{}{"apple", "banana", "orange", "grape", "peach"}
// 	keys := Rand(arr, 3)
// 	fmt.Println(keys) // [0 3 2]
//
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
