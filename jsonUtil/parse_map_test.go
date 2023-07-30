package jsonUtil

import (
	"reflect"
	"testing"
)

func TestParseMap(t *testing.T) {
	type TestData struct {
		Foo string
		Bar int
	}

	// 包含映射类型的测试用例
	testData := map[string]any{
		"Foo": "hello",
		"Bar": 42,
	}

	var result map[string]any
	err := parseMap(reflect.ValueOf(&result).Elem(), testData)
	if err != nil {
		t.Errorf("parseMap 失败: %s", err)
	}

	// 检查值是否被正确解析
	expectedResult := map[string]any{"Foo": "hello", "Bar": 42}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("parseMap 结果不匹配:\n期望值: %v\n实际值: %v", expectedResult, result)
	}
}

func TestParseValue(t *testing.T) {
	// 测试解析原始类型
	var intValue int
	err := parseValue(reflect.ValueOf(&intValue).Elem(), 42)
	if err != nil {
		t.Errorf("parseValue 失败: %s", err)
	}
	if intValue != 42 {
		t.Errorf("parseValue 结果不匹配: 期望值 42，实际值 %d", intValue)
	}

	// 测试解析结构体
	type TestStruct struct {
		Name string
		Age  int
	}

	var structValue TestStruct
	err = parseValue(reflect.ValueOf(&structValue).Elem(), map[string]any{"Name": "John", "Age": 30})
	if err != nil {
		t.Errorf("parseValue 失败: %s", err)
	}

	expectedStructValue := TestStruct{"John", 30}
	if !reflect.DeepEqual(structValue, expectedStructValue) {
		t.Errorf("parseValue 结果不匹配:\n期望值: %v\n实际值: %v", expectedStructValue, structValue)
	}

	// 测试解析切片
	var sliceValue []int
	err = parseValue(reflect.ValueOf(&sliceValue).Elem(), []any{1, 2, 3})
	if err != nil {
		t.Errorf("parseValue 失败: %s", err)
	}

	expectedSliceValue := []int{1, 2, 3}
	if !reflect.DeepEqual(sliceValue, expectedSliceValue) {
		t.Errorf("parseValue 结果不匹配:\n期望值: %v\n实际值: %v", expectedSliceValue, sliceValue)
	}

	// 测试不支持的类型
	var unsupportedValue complex64
	err = parseValue(reflect.ValueOf(&unsupportedValue).Elem(), 3.14)
	if err == nil {
		t.Errorf("parseValue 对于不支持的类型应该失败")
	}
}
