package sliceUtil

// Diff 计算差集
func Diff[T comparable](s []T, slices ...[]T) []T {
	seen := make(map[T]bool)
	for _, slice := range slices {
		for _, elem := range slice {
			seen[elem] = true
		}
	}

	var result []T
	for _, elem := range s {
		if !seen[elem] {
			result = append(result, elem)
		}
	}

	return result
}

// SymmetricDiff 计算对称差集
func SymmetricDiff[T comparable](slices ...[]T) []T {
	// 判断当前元素是否存在于其他切片中
	containsAllNotMe := func(item T, idx int, slices ...[]T) bool {
		for i, slice := range slices {
			if i == idx {
				continue
			}
			for _, val := range slice {
				if val == item {
					return true
				}
			}
		}
		return false
	}

	var result []T
	for idx, slice := range slices {
		for _, item := range slice {
			if !containsAllNotMe(item, idx, slices...) {
				result = append(result, item)
			}
		}
	}

	return UniqueSlice(result)
}

// Intersect 计算交集
func Intersect[T comparable](slices ...[]T) []T {
	// 判断元素是否存在于其他切片中
	containsAll := func(item T, slices [][]T) bool {
		for _, slice := range slices {
			if !InSlice(item, slice) {
				return false
			}
		}
		return true
	}

	var result []T
	for _, item := range slices[0] {
		if containsAll(item, slices[1:]) {
			result = append(result, item)
		}
	}
	return UniqueSlice(result)
}
