//go:build amd64 || arm64 || mips64 || mips64le || ppc64 || ppc64le || riscv64 || s390x || wasm

package anyx

// intFromInt64 将 int64 转换为 int。
// 上述架构上 int 与 int64 等宽，直接转换不会丢失精度，
// 因此无需范围检查（避免引入恒不触发的死分支影响覆盖率）。
//
// 注意：本文件的构建约束中每个架构都必须能被 CodeQL 识别为 64 位
// （见 int_size_convert.go 的拆分说明），loong64 虽然也是 64 位 int，
// 但不在 CodeQL 识别的架构列表中，故走 int_size_convert.go 的带检查版本。
func intFromInt64(v int64) (int, error) {
	return int(v), nil
}

// uintFromUint64 将 uint64 转换为 uint，说明同 intFromInt64。
func uintFromUint64(v uint64) (uint, error) {
	return uint(v), nil
}
