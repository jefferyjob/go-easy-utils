package sliceUtil

import (
	"reflect"
	"testing"
)

func TestExtractKeysInline(t *testing.T) {
	// 使用匿名结构体定义测试数据
	persons := []struct {
		ID   int
		Name string
	}{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
	}

	// 使用函数字面量定义 keyFunc 函数
	keys := ExtractKeys(persons, func(p struct {
		ID   int
		Name string
	}) int {
		return p.ID
	})

	expectedKeys := []int{1, 2, 3}
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, but got %v", expectedKeys, keys)
	}

	// 另一个例子：使用不同的结构体和 keyFunc 函数
	animals := []struct {
		ID   string
		Name string
	}{
		{"cat001", "Cat"},
		{"dog002", "Dog"},
	}

	animalKeys := ExtractKeys(animals, func(a struct {
		ID   string
		Name string
	}) string {
		return a.Name
	})

	expectedAnimalKeys := []string{"Cat", "Dog"}
	if !reflect.DeepEqual(animalKeys, expectedAnimalKeys) {
		t.Errorf("Expected animal keys %v, but got %v", expectedAnimalKeys, animalKeys)
	}
}
