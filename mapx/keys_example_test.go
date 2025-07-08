package mapx

import (
	"fmt"
	"sort"
)

func ExampleKeys() {
	m := map[string]int{"apple": 5, "banana": 3, "orange": 7}
	keys := Keys(m)

	// 对结果排序以确保输出一致（map 迭代顺序随机）
	sort.Strings(keys)
	fmt.Println("Keys:", keys)

	// Output: Keys: [apple banana orange]
}
