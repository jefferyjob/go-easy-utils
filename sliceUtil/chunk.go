// 提供了一个名为Chunk的函数，用于将一个切片按指定大小分成多个切片。
//
// 函数签名如下：
//
// 	func Chunk(slice []interface{}, size int) [][]interface{}
//
// 参数slice是要分割的切片，参数size是每个小切片的大小。
//
// 返回值是一个包含多个小切片的二维切片。如果切片长度不是size的整数倍，
// 则最后一个小切片可能会比指定的大小小。
//
// 示例：
//
// 	arr := []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
// 	chunks := Chunk(arr, 3)
// 	fmt.Println(chunks) // [[1 2 3] [4 5 6] [7 8]]
//
package sliceUtil

// Chunk 将一个切片按指定大小分成多个切片，并返回一个包含这些切片的二维切片。
//
// 参数slice是要分割的切片，参数size是每个小切片的大小。
//
// 返回值是一个包含多个小切片的二维切片。如果切片长度不是size的整数倍，
// 则最后一个小切片可能会比指定的大小小。
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
