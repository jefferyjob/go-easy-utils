package sliceUtil

import (
	"reflect"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	// Test case 1: Test with a slice of Person structs
	type Person struct {
		ID   int
		Name string
	}

	people := []Person{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	keyFunc := func(p Person) int {
		return p.ID
	}

	expected := map[int]Person{
		1: {ID: 1, Name: "Alice"},
		2: {ID: 2, Name: "Bob"},
		3: {ID: 3, Name: "Charlie"},
	}

	result := SliceToMap(people, keyFunc)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected, result)
	}

	// Test case 2: Test with an empty slice
	emptySlice := []Person{}
	expectedEmpty := map[int]Person{}

	resultEmpty := SliceToMap(emptySlice, keyFunc)

	if !reflect.DeepEqual(resultEmpty, expectedEmpty) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expectedEmpty, resultEmpty)
	}

	// Test case 3: Test with a slice containing elements with the same key
	duplicatePeople := []Person{
		{ID: 1, Name: "Alice"},
		{ID: 1, Name: "Bob"},
	}

	expectedDuplicate := map[int]Person{
		1: {ID: 1, Name: "Bob"}, // The last occurrence of the duplicate key will overwrite the previous one
	}

	resultDuplicate := SliceToMap(duplicatePeople, keyFunc)

	if !reflect.DeepEqual(resultDuplicate, expectedDuplicate) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expectedDuplicate, resultDuplicate)
	}
}
