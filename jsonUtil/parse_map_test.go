package jsonUtil

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestParseMap(t *testing.T) {
	tests := []struct {
		name       string
		inputValue func() reflect.Value
		inputData  func() map[string]any
		wantErr    error
	}{
		{
			name: "正常的map解析",
			inputValue: func() reflect.Value {
				var result map[string]any
				return reflect.ValueOf(&result).Elem()
			},
			inputData: func() map[string]any {
				jsonData := `{
					"Foo":"hello",
					"Bar":42
				}`
				var data map[string]any
				err := json.Unmarshal([]byte(jsonData), &data)
				require.NoError(t, err)

				return map[string]any{
					"Foo": "hello",
					"Bar": 42,
				}
			},
		},
		//{
		//	name: "正常的map解析",
		//	inputValue: func() reflect.Value {
		//		var result map[string]any
		//		return reflect.ValueOf(&result).Elem()
		//	},
		//	inputData: map[string]any{
		//		"Foo": "hello",
		//		"Bar": 42,
		//	},
		//},
		//{
		//	name: "非Map数据类型",
		//	inputValue: func() reflect.Value {
		//		return reflect.ValueOf(42)
		//	},
		//	inputData: map[string]any{"Foo": "hello", "Bar": 42},
		//	wantErr:   ErrNotMap,
		//},
		//{
		//	name: "Map嵌套",
		//	inputValue: func() reflect.Value {
		//		return reflect.New(reflect.TypeOf(map[string]any{})).Elem()
		//	},
		//	inputData: map[string]any{
		//		"nested": map[string]any{
		//			"subKey": "subValue",
		//		},
		//	},
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parseMap(tt.inputValue(), tt.inputData())
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestParseMap3(t *testing.T) {
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

	// 测试解析接口
	var interfaceValue interface{}
	err = parseValue(reflect.ValueOf(&interfaceValue).Elem(), "test")
	if err != nil {
		t.Errorf("parseValue 失败: %s", err)
	}

	expectedInterfaceValue := "test"
	if !reflect.DeepEqual(interfaceValue, expectedInterfaceValue) {
		t.Errorf("parseValue 结果不匹配:\n期望值: %v\n实际值: %v", expectedInterfaceValue, interfaceValue)
	}

	// 测试不支持的类型
	var unsupportedValue complex64
	err = parseValue(reflect.ValueOf(&unsupportedValue).Elem(), 3.14)
	if err == nil {
		t.Errorf("parseValue 对于不支持的类型应该失败")
	}
}

func TestParseMap2(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]any
		output    reflect.Value
		expectErr error
	}{
		{
			name:      "Non-Map Value",
			input:     map[string]any{"key": "value"},
			output:    reflect.ValueOf("non-map value"),
			expectErr: ErrNotMap,
		},
		{
			name:      "Map with Primitive Values",
			input:     map[string]any{"key1": "value1", "key2": 123},
			output:    reflect.New(reflect.TypeOf(map[string]any{})).Elem(),
			expectErr: nil,
		},
		{
			name: "Map with Nested Map",
			input: map[string]any{
				"nested": map[string]any{
					"subKey": "subValue",
				},
			},
			output:    reflect.New(reflect.TypeOf(map[string]any{})).Elem(),
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parseMap(tt.output, tt.input)
			if !errors.Is(err, tt.expectErr) {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}
		})
	}
}
