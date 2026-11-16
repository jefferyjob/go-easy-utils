package anyx

// intFromInt64 将 int64 转换为 int。
// amd64 架构上 int 与 int64 等宽，直接转换不会丢失精度，
// 因此无需范围检查（避免引入恒不触发的死分支影响覆盖率）。
//
// 说明：本文件通过文件名后缀 _amd64 隐式约束构建架构（Go 标准惯例）。
// CodeQL 的 go/incorrect-integer-conversion 查询仅能通过文件名隐式约束
// 识别 int 位宽，显式的 //go:build amd64 || arm64 || ... 形式不会被识别，
// 会将直通转换误报为缺少上界检查，因此按架构拆分为独立文件。
func intFromInt64(v int64) (int, error) {
	return int(v), nil
}

// uintFromUint64 将 uint64 转换为 uint，说明同 intFromInt64。
func uintFromUint64(v uint64) (uint, error) {
	return uint(v), nil
}
