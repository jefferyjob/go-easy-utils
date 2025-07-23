package anyx

import (
	"testing"
)

func TestTernary(t *testing.T) {
	tests := []struct {
		name     string
		expr     bool
		a        any
		b        any
		expected any
	}{
		{"TrueCase_String", true, "yes", "no", "yes"},
		{"FalseCase_String", false, "yes", "no", "no"},
		{"TrueCase_Int", true, 1, 2, 1},
		{"FalseCase_Int", false, 1, 2, 2},
		{"TrueCase_Struct", true, struct{ X int }{1}, struct{ X int }{2}, struct{ X int }{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.a.(type) {
			case string:
				result := Ternary(tt.expr, tt.a.(string), tt.b.(string))
				if result != tt.expected.(string) {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			case int:
				result := Ternary(tt.expr, tt.a.(int), tt.b.(int))
				if result != tt.expected.(int) {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			case struct{ X int }:
				result := Ternary(tt.expr, tt.a.(struct{ X int }), tt.b.(struct{ X int }))
				if result != tt.expected.(struct{ X int }) {
					t.Errorf("expected %+v, got %+v", tt.expected, result)
				}
			}
		})
	}
}
