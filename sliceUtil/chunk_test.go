package sliceUtil

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		name  string
		slice []interface{}
		size  int
		want  [][]interface{}
	}{
		{
			name:  "empty slice",
			slice: []interface{}{},
			size:  3,
			want:  [][]interface{}{},
		},
		{
			name:  "slice with less than chunk size",
			slice: []interface{}{1, 2},
			size:  3,
			want:  [][]interface{}{{1, 2}},
		},
		{
			name:  "slice with exact chunk size",
			slice: []interface{}{1, 2, 3, 4, 5, 6},
			size:  3,
			want:  [][]interface{}{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name:  "slice with more than chunk size",
			slice: []interface{}{1, 2, 3, 4, 5, 6, 7},
			size:  3,
			want:  [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.slice, tt.size)
			if len(got) != len(tt.want) {
				t.Errorf("Chunk() = %+v, want %+v", got, tt.want)
				return
			}
			for i := range got {
				if !reflect.DeepEqual(got[i], tt.want[i]) {
					t.Errorf("Chunk() = %+v, want %+v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestChunkStr(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	size := 3
	expected := [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k"}}
	result := Chunk(slice, size)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkStr(%v, %d) = %v; expected %v", slice, size, result, expected)
	}
}

func TestChunkInt(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	size := 4
	expected := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	result := Chunk(slice, size)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkInt(%v, %d) = %v; expected %v", slice, size, result, expected)
	}
}

func TestChunkInt8(t *testing.T) {
	slice := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	size := 4
	expected := [][]int8{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	result := Chunk(slice, size)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkInt8(%v, %d) = %v; expected %v", slice, size, result, expected)
	}
}

func TestChunkInt32(t *testing.T) {
	slice := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	size := 4
	expected := [][]int32{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	result := Chunk(slice, size)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkInt32(%v, %d) = %v; expected %v", slice, size, result, expected)
	}
}

func TestChunkInt64(t *testing.T) {
	slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	size := 4
	expected := [][]int64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	result := Chunk(slice, size)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkInt64(%v, %d) = %v; expected %v", slice, size, result, expected)
	}
}

func TestChunkUint(t *testing.T) {
	slice := []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	size := 4
	expected := [][]uint{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	result := Chunk(slice, size)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkUint(%v, %d) = %v; expected %v", slice, size, result, expected)
	}
}

func TestChunkUint8(t *testing.T) {
	slice := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	size := 4
	expected := [][]uint8{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	result := Chunk(slice, size)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkUint8(%v, %d) = %v; expected %v", slice, size, result, expected)
	}
}

func TestChunkUint16(t *testing.T) {
	slice := []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	size := 4
	expected := [][]uint16{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	result := Chunk(slice, size)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkUint16(%v, %d) = %v; expected %v", slice, size, result, expected)
	}
}

func TestChunkUint32(t *testing.T) {
	slice := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	size := 4
	expected := [][]uint32{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	result := Chunk(slice, size)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkUint32(%v, %d) = %v; expected %v", slice, size, result, expected)
	}
}

func TestChunkUint64(t *testing.T) {
	slice := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	size := 4
	expected := [][]uint64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}}
	result := Chunk(slice, size)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ChunkUint64(%v, %d) = %v; expected %v", slice, size, result, expected)
	}
}