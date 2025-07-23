package anyx

// Ternary 三元运算符
// 根据表达式 expr 的布尔值，返回 a 或 b 中的一个
// 如果 expr 为 true，则返回 a；否则返回 b
func Ternary[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}
