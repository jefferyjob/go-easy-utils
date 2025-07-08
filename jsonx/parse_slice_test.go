package jsonx

import (
	"reflect"
	"testing"
)

type ExampleStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestParseSlice(t *testing.T) {
	var slice []any

	// 添加元素到切片
	slice = append(slice, "hello")
	slice = append(slice, "world")

	// 创建一个新的切片值
	sliceValue := reflect.ValueOf(&[]string{}).Elem()

	// 解析切片
	err := parseSlice(sliceValue, slice)
	if err != nil {
		t.Errorf("expected no error, but got %s", err.Error())
	}

	// 检查解析的切片是否与预期相等
	expectedSlice := []string{"hello", "world"}
	actualSlice := sliceValue.Interface().([]string)
	if !reflect.DeepEqual(expectedSlice, actualSlice) {
		t.Errorf("expected %v, but got %v", expectedSlice, actualSlice)
	}

	// 测试解析结构体切片
	var structSlice []any
	structSlice = append(structSlice, ExampleStruct{"Tom", 30})
	structSlice = append(structSlice, ExampleStruct{"Jerry", 25})

	// 创建一个新的切片值
	structSliceValue := reflect.ValueOf(&[]ExampleStruct{}).Elem()

	// 解析切片
	err = parseSlice(structSliceValue, structSlice)
	if err != nil {
		t.Errorf("expected no error, but got %s", err.Error())
	}

	// 检查解析的切片是否与预期相等
	expectedStructSlice := []ExampleStruct{{"Tom", 30}, {"Jerry", 25}}
	actualStructSlice := structSliceValue.Interface().([]ExampleStruct)
	if !reflect.DeepEqual(expectedStructSlice, actualStructSlice) {
		t.Errorf("expected %v, but got %v", expectedStructSlice, actualStructSlice)
	}
}
