package jsonx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToStructInvalidJSON(t *testing.T) {
	testCases := []struct {
		name    string
		jsonStr string
	}{
		{
			name:    "空字符串",
			jsonStr: "",
		},
		{
			name:    "非法JSON",
			jsonStr: `{"name": }`,
		},
		{
			name:    "JSON数组",
			jsonStr: `[1, 2, 3]`,
		},
		{
			name:    "JSON字符串",
			jsonStr: `"hello"`,
		},
		{
			name:    "JSON数字",
			jsonStr: `123`,
		},
		{
			name:    "JSON布尔",
			jsonStr: `true`,
		},
		{
			name:    "尾随垃圾数据",
			jsonStr: `{"name": "Bob"} x`,
		},
		{
			name:    "尾随第二个JSON值",
			jsonStr: `{"name": "Bob"} {"name": "Alice"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var s tsS1
			err := ToStruct(tc.jsonStr, &s)
			require.Error(t, err)
		})
	}
}

func TestToStructNullJSON(t *testing.T) {
	var s tsS1
	err := ToStruct(`null`, &s)
	require.NoError(t, err)
	require.Equal(t, tsS1{}, s)
}

func TestToStructTrailingData(t *testing.T) {
	// 尾随一个合法的完整 JSON 值时返回明确错误，并可用 ErrSyntax 判断
	var s tsS1
	err := ToStruct(`{"name": "Bob"} []`, &s)
	require.ErrorIs(t, err, ErrSyntax)
	require.Contains(t, err.Error(), "unexpected data after top-level JSON value")
}

func TestToStructNotPointer(t *testing.T) {
	t.Run("非指针类型", func(t *testing.T) {
		s := tsS1{}
		err := ToStruct(tsJ1, s)
		require.ErrorIs(t, err, ErrPoint)
	})

	t.Run("指针指向非结构体", func(t *testing.T) {
		num := 10
		err := ToStruct(tsJ1, &num)
		require.ErrorIs(t, err, ErrPoint)
	})

	t.Run("nil参数", func(t *testing.T) {
		err := ToStruct(tsJ1, nil)
		require.ErrorIs(t, err, ErrPoint)
	})

	t.Run("nil结构体指针", func(t *testing.T) {
		var s *tsS1
		err := ToStruct(tsJ1, s)
		require.ErrorIs(t, err, ErrPoint)
	})
}
