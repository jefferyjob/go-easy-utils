package mapUtil

import (
	"fmt"
	"sort"
)

func ExampleExtractValues() {
	m := map[string]int{"apple": 5, "banana": 3, "orange": 7}
	values := ExtractValues(m)

	// 对结果排序以确保输出一致（map 迭代顺序随机）
	sort.Ints(values)
	fmt.Println("Values:", values)

	// Output: Values: [3 5 7]
}
