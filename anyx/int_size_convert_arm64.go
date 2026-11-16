package anyx

// intFromInt64 将 int64 转换为 int。
// arm64 架构上 int 与 int64 等宽，直接转换不会丢失精度，
// 因此无需范围检查（避免引入恒不触发的死分支影响覆盖率）。
// 文件名隐式约束的说明同 int_size_convert_amd64.go。
func intFromInt64(v int64) (int, error) {
	return int(v), nil
}

// uintFromUint64 将 uint64 转换为 uint，说明同 intFromInt64。
func uintFromUint64(v uint64) (uint, error) {
	return uint(v), nil
}
