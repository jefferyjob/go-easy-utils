package jsonUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	var iPar = "point"
	tests := []struct {
		name  string
		value any
		want  string
	}{
		{"nil", nil, ""},
		{"string", "hello", "hello"},
		{"int", 42, "42"},
		{"int8", int8(42), "42"},
		{"int16", int16(42), "42"},
		{"int32", int32(42), "42"},
		{"int64", int64(42), "42"},
		{"uint", uint(42), "42"},
		{"uint8", uint8(42), "42"},
		{"uint16", uint16(42), "42"},
		{"uint32", uint32(42), "42"},
		{"uint64", uint64(42), "42"},
		{"float32", float32(3.14159), "3.14159"},
		{"float64", 3.14159, "3.14159"},
		{"bool-true", true, "true"},
		{"bool-false", false, "false"},
		{"point", &iPar, "point"},
		{"complex64", complex64(1 + 2i), "(1+2i)"},
		{"complex128", complex128(3 + 4i), "(3+4i)"},
		{"chan", make(chan int), ""},
		{"int nil", (*int)(nil), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toStringReflect(tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}
