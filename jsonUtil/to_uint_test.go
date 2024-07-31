package jsonUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToUint64(t *testing.T) {
	var iPtr = 90
	tests := []struct {
		name      string
		input     any
		want      uint64
		wantError error
	}{
		{
			name:  "Test -float32",
			input: float32(-0.1),
			want:  0,
		},
		{
			name:  "Test -float64",
			input: float64(-0.2),
			want:  0,
		},
		{
			name:  "Test -int",
			input: int(-1),
			want:  0,
		},
		{
			name:  "Test -int8",
			input: int8(-2),
			want:  0,
		},
		{
			name:  "Test -int16",
			input: int16(-3),
			want:  0,
		},
		{
			name:  "Test -int32",
			input: int32(-4),
			want:  0,
		},
		{
			name:  "Test -int64",
			input: int64(-5),
			want:  0,
		},
		{
			name:  "Test uint",
			input: uint(12),
			want:  12,
		},
		{
			name:  "Test uint8",
			input: uint8(42),
			want:  42,
		},
		{
			name:  "Test uint16",
			input: uint16(42),
			want:  42,
		},
		{
			name:  "Test uint32",
			input: uint32(42),
			want:  42,
		},
		{
			name:  "Test uint64",
			input: uint64(42),
			want:  42,
		},
		{
			name:  "Test int8",
			input: int8(42),
			want:  42,
		},
		{
			name:  "Test int16",
			input: int16(42),
			want:  42,
		},
		{
			name:  "Test int32",
			input: int32(42),
			want:  42,
		},
		{
			name:  "Test int64",
			input: int64(42),
			want:  42,
		},
		{
			name:  "Test float32",
			input: float32(42.0),
			want:  42,
		},
		{
			name:  "Test float64",
			input: float64(42.0),
			want:  42,
		},
		{
			name:  "Test complex64",
			input: complex64(complex(42, 0)),
			want:  42,
		},
		{
			name:  "Test complex128",
			input: complex128(complex(42, 0)),
			want:  42,
		},
		{
			name:  "test -complex",
			input: complex(-1, -1),
			want:  0,
		},
		{
			name:  "Test string",
			input: "42",
			want:  42,
		},
		{
			name:      "Test invalid string",
			input:     "not a number",
			wantError: ErrSyntax,
		},
		{
			name:      "Test nil pointer",
			input:     (*int)(nil),
			want:      0,
			wantError: nil,
		},
		{
			name:  "Test bool true",
			input: true,
			want:  1,
		},
		{
			name:  "Test bool false",
			input: false,
			want:  0,
		},
		{
			name:      "Test invalid type",
			input:     make(chan int),
			wantError: ErrType,
		},
		{
			name:  "Test point",
			input: &iPtr,
			want:  90,
		},
		{
			name:  "nil",
			input: nil,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toUint64Reflect(tt.input)
			assert.Equal(t, tt.wantError, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
