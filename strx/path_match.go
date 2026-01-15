package strx

import "strings"

// PathMatch 判断权限路径模式是否可以匹配实际请求路径
//
// 支持的通配规则：
//   - *   ：单层路径匹配（仅匹配一个 segment）
//   - **  ：多层路径匹配（必须至少匹配一层）
//   - users* / *Test ：同一层级内的段内通配，不可跨越 '/'
//
// 该方法会统一对路径进行标准化处理，然后按路径段（segment）逐层匹配。
// 适用于权限校验场景，匹配规则严格、行为可预测。
func PathMatch(pattern, path string) bool {
	pattern = normalizePath(pattern)
	path = normalizePath(path)

	pSegs := strings.Split(pattern, "/")
	rSegs := strings.Split(path, "/")

	return matchSegments(pSegs, rSegs)
}

// normalizePath 对路径进行标准化：
//   - 确保以 '/' 开头
//   - 移除末尾多余的 '/'
//
// 目的是避免因路径书写格式差异（如是否带尾部 /）导致匹配结果不一致。
func normalizePath(p string) string {
	if p == "" {
		return "/"
	}
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	return strings.TrimSuffix(p, "/")
}

// matchSegments 按路径段（以 '/' 分隔）逐段进行匹配。
//
// 匹配规则说明：
//   - 普通段：必须与请求段匹配（支持段内 * 通配）
//   - *      ：匹配任意单个路径段
//   - **     ：匹配多个路径段，但至少匹配一层
//
// 该函数是核心匹配逻辑，严格保证权限路径的匹配边界，
// 不允许 ** 匹配 0 层，避免权限过宽。
func matchSegments(pSegs, rSegs []string) bool {
	i, j := 0, 0

	for i < len(pSegs) && j < len(rSegs) {
		switch pSegs[i] {

		case "**":
			// ** 必须至少匹配一层
			if i+1 == len(pSegs) {
				return j < len(rSegs)
			}

			// 吃掉至少一层，然后尝试匹配后续
			for k := j + 1; k <= len(rSegs); k++ {
				if matchSegments(pSegs[i+1:], rSegs[k:]) {
					return true
				}
			}
			return false

		case "*":
			// 单层匹配
			i++
			j++

		default:
			// 普通 segment（支持 users* 这种）
			if !matchSegment(pSegs[i], rSegs[j]) {
				return false
			}
			i++
			j++
		}
	}

	// pattern 和 path 必须同时结束
	if i == len(pSegs) && j == len(rSegs) {
		return true
	}

	// pattern 剩下的只能是 **，但 ** 不能匹配 0 层
	if i == len(pSegs)-1 && pSegs[i] == "**" && j < len(rSegs) {
		return true
	}

	return false
}

// matchSegment 判断单个路径段是否匹配。
//
// 支持两种情况：
//   - 完全相等
//   - 段内通配（如 users*、*Test）
//
// 注意：
//   - 该匹配仅作用于单个 segment
//   - 不允许跨越 '/'，以保证路径层级的严格性
func matchSegment(pattern, segment string) bool {
	// 完全相等
	if pattern == segment {
		return true
	}

	// 段内通配：users*
	if strings.Contains(pattern, "*") {
		return matchWithStar(pattern, segment)
	}

	return false
}

// matchWithStar 实现单个路径段内的 '*' 通配匹配。
//
// 规则说明：
//   - 仅支持一个 '*'
//   - '*' 可匹配任意长度字符，但不能跨越路径分隔符
//   - 适用于 users*、*Service、user*Test 等规则
func matchWithStar(pattern, str string) bool {
	// 仅支持一个 *
	idx := strings.Index(pattern, "*")
	if idx == -1 {
		return pattern == str
	}

	prefix := pattern[:idx]
	suffix := pattern[idx+1:]

	if !strings.HasPrefix(str, prefix) {
		return false
	}
	if suffix != "" && !strings.HasSuffix(str, suffix) {
		return false
	}

	return len(str) >= len(prefix)+len(suffix)
}
