package validUtil

import "testing"

func TestIsPostalCode(t *testing.T) {
	// Test valid postal codes
	validCodes := []string{"100000", "200000", "999999"}
	for _, code := range validCodes {
		if !IsPostalCode(code) {
			t.Errorf("%s should be a valid postal code", code)
		}
	}

	// Test invalid postal codes
	invalidCodes := []string{"1234567", "2000000", "123a56"}
	for _, code := range invalidCodes {
		if IsPostalCode(code) {
			t.Errorf("%s should be an invalid postal code", code)
		}
	}
}
