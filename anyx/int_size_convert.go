//go:build !(amd64 || arm64)

package anyx

import "math"

// intFromInt64 将 int64 安全地转换为 int。
// int 在 32 位平台上为 32 位，直接转换会静默截断高位，因此这里检查范围，
// 超出范围时返回 ErrValOut；在 64 位平台上 int 与 int64 等宽，
// math.MinInt/math.MaxInt 随编译目标自适应，检查恒不触发。
// 本文件适用于除 amd64/arm64 外的所有架构（含 32 位平台、loong64 及未来新架构）。
func intFromInt64(v int64) (int, error) {
	if v < math.MinInt || v > math.MaxInt {
		return 0, ErrValOut
	}
	return int(v), nil
}

// uintFromUint64 将 uint64 安全地转换为 uint。
// uint 在 32 位平台上为 32 位，直接转换会静默截断高位，因此这里检查范围，
// 超出范围时返回 ErrValOut；在 64 位平台上 uint 与 uint64 等宽，
// math.MaxUint 随编译目标自适应，检查恒不触发。
func uintFromUint64(v uint64) (uint, error) {
	if v > math.MaxUint {
		return 0, ErrValOut
	}
	return uint(v), nil
}
