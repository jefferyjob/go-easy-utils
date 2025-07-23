package strx

import "strings"

// Cut 删除 s 中出现的 sub 字符串
// 第三个参数 n 可选：
//   - 省略    → 删除所有匹配（等价 strings.ReplaceAll）
//   - n == 1  → 只删除第一次出现
//   - n > 1   → 删除前 n 次出现
func Cut(s, sub string, n ...int) string {
	if len(sub) == 0 || len(s) == 0 {
		return s // nothing to do
	}

	// 默认: 删除所有
	count := -1
	if len(n) > 0 {
		count = n[0]
	}
	return strings.Replace(s, sub, "", count)
}
