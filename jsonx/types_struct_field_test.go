package jsonx

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToStructStructField(t *testing.T) {
	type s struct {
		Addr    Address `json:"addr"`
		Missing Address `json:"missing"`
	}

	var v s
	err := ToStruct(`{"addr": {"country": "China", "city": "Shanghai", "postcode": 200000}}`, &v)
	require.NoError(t, err)
	require.Equal(t, "China", v.Addr.Country)
	require.Equal(t, "Shanghai", v.Addr.City)
	require.Equal(t, 200000, v.Addr.Postcode)
	// JSON 中不存在的字段保持零值
	require.Equal(t, Address{}, v.Missing)
}

func TestToStructStructFieldError(t *testing.T) {
	type s struct {
		Addr Address `json:"addr"`
	}

	var v s
	err := ToStruct(`{"addr": "Shanghai"}`, &v)
	require.Error(t, err)
	require.Contains(t, err.Error(), "cannot unmarshal string into")
}

func TestToStructStructFieldNull(t *testing.T) {
	type s struct {
		Addr Address `json:"addr"`
	}

	var v s
	err := ToStruct(`{"addr": null}`, &v)
	require.NoError(t, err)
	require.Equal(t, Address{}, v.Addr)
}

func TestToStructAnyField(t *testing.T) {
	type s struct {
		Nil  any `json:"nil"`
		Num  any `json:"num"`
		Str  any `json:"str"`
		Bool any `json:"bool"`
		Arr  any `json:"arr"`
		Obj  any `json:"obj"`
	}

	var v s
	err := ToStruct(`{"nil": null, "num": 1, "str": "x", "bool": true, "arr": [1, "a"], "obj": {"k": "v"}}`, &v)
	require.NoError(t, err)
	require.Nil(t, v.Nil)
	// UseNumber 后 any 字段中的数字保留为 json.Number 字面量
	require.Equal(t, json.Number("1"), v.Num)
	require.Equal(t, "x", v.Str)
	require.Equal(t, true, v.Bool)
	require.Equal(t, []any{json.Number("1"), "a"}, v.Arr)
	require.Equal(t, map[string]any{"k": "v"}, v.Obj)
}

func TestToStructUnsupportedField(t *testing.T) {
	type s struct {
		C complex128 `json:"c"`
	}

	var v s
	err := ToStruct(`{"c": 1}`, &v)
	require.ErrorIs(t, err, ErrUnsupported)
	require.Contains(t, err.Error(), "不支持类型 complex128")
}

func TestToStructSkipFields(t *testing.T) {
	type s struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		city    string // 未导出字段不可设置，跳过
		NoTag   string // 无 json tag，跳过
		Ignored string `json:"-"` // "-" 表示忽略该字段，即使 JSON 中有 "-" key 也不填充
	}

	var v s
	err := ToStruct(`{"name": "Bob", "age": 30, "city": "Shanghai", "NoTag": "x", "-": "y", "extra": "ignored"}`, &v)
	require.NoError(t, err)
	require.Equal(t, "Bob", v.Name)
	require.Equal(t, 30, v.Age)
	require.Equal(t, "", v.city)
	require.Equal(t, "", v.NoTag)
	require.Equal(t, "", v.Ignored)
}

func TestToStructTagWithOptions(t *testing.T) {
	type s struct {
		A int    `json:"a,omitempty"` // 带选项的 tag 应匹配 "a" 而不是 "a,omitempty"
		B int    `json:"b,string"`    // string 选项同理
		C string `json:",omitempty"`  // 逗号前名称为空，与无 tag 字段一致，仅 "" key 可命中
	}

	var v s
	err := ToStruct(`{"a": 1, "b": "42"}`, &v)
	require.NoError(t, err)
	require.Equal(t, 1, v.A)
	require.Equal(t, 42, v.B)
	require.Equal(t, "", v.C)
}

func TestToStructInterfaceFieldNotAssignable(t *testing.T) {
	type stringer interface {
		String() string
	}
	type s struct {
		V stringer `json:"v"`
	}

	var v s
	err := ToStruct(`{"v": "x"}`, &v)
	require.ErrorIs(t, err, ErrUnsupported)
	require.Contains(t, err.Error(), "无法赋值给接口类型")
}

func TestToStructNestedErrorPath(t *testing.T) {
	type inner struct {
		Nums []int `json:"nums"`
	}
	type outer struct {
		In inner `json:"in"`
	}

	var v outer
	err := ToStruct(`{"in": {"nums": ["1", "x"]}}`, &v)
	require.ErrorIs(t, err, ErrSyntax)
	require.Contains(t, err.Error(), "in.nums[1]")
}
