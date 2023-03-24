package floatUtil

import (
	"math"
	"testing"
)

func TestFloat32ToStr(t *testing.T) {
	var f float32 = 3.14159
	want := "3.14159"
	got := Float32ToStr(f)
	if got != want {
		t.Errorf("Float32ToStr(%v) = %v; want %v", f, got, want)
	}
}

func TestFloat64ToStr(t *testing.T) {
	var f float64 = 3.14159
	want := "3.14159"
	got := Float64ToStr(f)
	if got != want {
		t.Errorf("Float64ToStr(%v) = %v; want %v", f, got, want)
	}
}

func TestFloat32ToFloat64(t *testing.T) {
	var f float32 = 3.14159
	want := 3.14159
	got := Float32ToFloat64(f)
	if math.Abs(got-want) > 1e-6 {
		t.Errorf("Float32ToFloat64(%v) = %v; want %v", f, got, want)
	}
}

func TestFloat64ToFloat32(t *testing.T) {
	var f float64 = 3.14159
	want := float32(3.14159)
	got := Float64ToFloat32(f)
	if math.Abs(float64(got)-float64(want)) > 1e-6 {
		t.Errorf("Float64ToFloat32(%v) = %v; want %v", f, got, want)
	}
}
