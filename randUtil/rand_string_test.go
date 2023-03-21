package randUtil

import "testing"

func TestRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		got := RandomString(10)
		if len(got) != 10 {
			t.Errorf("RandomString(10) = %q, want length %d", got, 10)
		}
	}
}

func TestRandomBytes(t *testing.T) {
	for i := 0; i < 10; i++ {
		got := RandomBytes(10)
		if len(got) != 10 {
			t.Errorf("RandomBytes(10) = %v, want length %d", got, 10)
		}
	}
}
