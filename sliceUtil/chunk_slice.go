package sliceUtil

// Chunk 将一个切片按指定大小分成多个切片，并返回一个包含这些切片的二维切片。
func Chunk(slice []interface{}, size int) [][]interface{} {
	var chunks [][]interface{}
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ChunkStr 将一个 string 类型的切片切割成指定大小的多个子切片
func ChunkStr(slice []string, size int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ChunkInt 将一个 int 类型的切片切割成指定大小的多个子切片
func ChunkInt(slice []int, size int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ChunkInt8 将一个 int8 类型的切片切割成指定大小的多个子切片
func ChunkInt8(slice []int8, size int) [][]int8 {
	var chunks [][]int8
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ChunkInt32 将一个 int32 类型的切片切割成指定大小的多个子切片
func ChunkInt32(slice []int32, size int) [][]int32 {
	var chunks [][]int32
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ChunkInt64 将一个 int64 类型的切片切割成指定大小的多个子切片
func ChunkInt64(slice []int64, size int) [][]int64 {
	var chunks [][]int64
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ChunkUint 将一个 uint 类型的切片切割成指定大小的多个子切片
func ChunkUint(slice []uint, size int) [][]uint {
	var chunks [][]uint
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ChunkUint8 将一个 uint8 类型的切片切割成指定大小的多个子切片
func ChunkUint8(slice []uint8, size int) [][]uint8 {
	var chunks [][]uint8
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ChunkUint16 将一个 uint16 类型的切片切割成指定大小的多个子切片
func ChunkUint16(slice []uint16, size int) [][]uint16 {
	var chunks [][]uint16
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ChunkUint32 将一个 uint32 类型的切片切分成固定大小的子切片
func ChunkUint32(slice []uint32, size int) [][]uint32 {
	var chunks [][]uint32
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// ChunkUint64 将一个 uint64 类型的切片切分成固定大小的子切片
func ChunkUint64(slice []uint64, size int) [][]uint64 {
	var chunks [][]uint64
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
