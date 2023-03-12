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
