package mapUtil

import "testing"

func TestMapValueExists(t *testing.T) {
	m := map[string]string{
		"foo": "bar",
		"baz": "123",
	}
	if !MapValueExists(m, "bar") {
		t.Errorf("expected MapValueExists to return true for value 'bar'")
	}
	if !MapValueExists(m, "123") {
		t.Errorf("expected MapValueExists to return true for value 123")
	}
	if MapValueExists(m, "non-existent") {
		t.Errorf("expected MapValueExists to return false for value 'non-existent'")
	}
}

func TestMapKeyExists(t *testing.T) {
	m := map[string]any{
		"foo": "bar",
		"baz": 123,
		"qux": []int{1, 2, 3},
	}
	if !MapKeyExists(m, "foo") {
		t.Errorf("expected MapKeyExists to return true for key 'foo'")
	}
	if !MapKeyExists(m, "baz") {
		t.Errorf("expected MapKeyExists to return true for key 'baz'")
	}
	if MapKeyExists(m, "non-existent") {
		t.Errorf("expected MapKeyExists to return false for key 'non-existent'")
	}
}
