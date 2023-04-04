package jsonUtil

import "testing"

func TestToString(t *testing.T) {
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
		{"complex64", complex64(1 + 2i), "(1+2i)"},
		{"complex128", complex128(3 + 4i), "(3+4i)"},
		{"chan", make(chan int), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toString(tt.value); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}

			if got := toStringReflect(tt.value); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
