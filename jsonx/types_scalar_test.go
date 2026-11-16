package jsonx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToStructBool(t *testing.T) {
	type s struct {
		A bool `json:"a"`
		B bool `json:"b"`
		C bool `json:"c"`
	}

	var v s
	err := ToStruct(`{"a": true, "b": "false", "c": 1}`, &v)
	require.NoError(t, err)
	require.Equal(t, true, v.A)
	require.Equal(t, false, v.B)
	// 数字类型无法识别为 bool，保持零值
	require.Equal(t, false, v.C)
}

func TestToStructBoolError(t *testing.T) {
	type s struct {
		B bool `json:"b"`
	}

	var v s
	err := ToStruct(`{"b": "not-a-bool"}`, &v)
	require.Error(t, err)
	require.Contains(t, err.Error(), "cannot convert not-a-bool to bool")
}

func TestToStructInt(t *testing.T) {
	type s struct {
		A int   `json:"a"`
		B int8  `json:"b"`
		C int16 `json:"c"`
		D int32 `json:"d"`
		E int64 `json:"e"`
		P *int  `json:"p"`
	}

	var v s
	err := ToStruct(`{"a": 42, "b": "7", "c": -300, "d": "1024", "e": 9999999999, "p": "66"}`, &v)
	require.NoError(t, err)
	require.Equal(t, 42, v.A)
	require.Equal(t, int8(7), v.B)
	require.Equal(t, int16(-300), v.C)
	require.Equal(t, int32(1024), v.D)
	require.Equal(t, int64(9999999999), v.E)
	require.NotNil(t, v.P)
	require.Equal(t, 66, *v.P)
}

func TestToStructIntError(t *testing.T) {
	t.Run("字符串解析失败", func(t *testing.T) {
		type s struct {
			A int `json:"a"`
		}
		var v s
		err := ToStruct(`{"a": "abc"}`, &v)
		require.ErrorIs(t, err, ErrSyntax)
	})

	t.Run("不支持的数据类型", func(t *testing.T) {
		type s struct {
			A int `json:"a"`
		}
		var v s
		err := ToStruct(`{"a": true}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), "无法转换 true 为 int")
	})

	t.Run("指针字段错误传递", func(t *testing.T) {
		type s struct {
			A *int `json:"a"`
		}
		var v s
		err := ToStruct(`{"a": "abc"}`, &v)
		require.ErrorIs(t, err, ErrSyntax)
	})
}

func TestToStructUint(t *testing.T) {
	type s struct {
		A uint   `json:"a"`
		B uint8  `json:"b"`
		C uint16 `json:"c"`
		D uint32 `json:"d"`
		E uint64 `json:"e"`
		P *uint  `json:"p"`
	}

	var v s
	err := ToStruct(`{"a": 42, "b": "200", "c": 65535, "d": "4294967295", "e": "18446744073709551615", "p": 8}`, &v)
	require.NoError(t, err)
	require.Equal(t, uint(42), v.A)
	require.Equal(t, uint8(200), v.B)
	require.Equal(t, uint16(65535), v.C)
	require.Equal(t, uint32(4294967295), v.D)
	require.Equal(t, uint64(18446744073709551615), v.E)
	require.NotNil(t, v.P)
	require.Equal(t, uint(8), *v.P)
}

func TestToStructUintError(t *testing.T) {
	t.Run("字符串解析失败", func(t *testing.T) {
		type s struct {
			A uint `json:"a"`
		}
		var v s
		err := ToStruct(`{"a": "abc"}`, &v)
		require.ErrorIs(t, err, ErrSyntax)
	})

	t.Run("负数字符串", func(t *testing.T) {
		type s struct {
			A uint `json:"a"`
		}
		var v s
		err := ToStruct(`{"a": "-1"}`, &v)
		require.ErrorIs(t, err, ErrSyntax)
	})

	t.Run("不支持的数据类型", func(t *testing.T) {
		type s struct {
			A uint `json:"a"`
		}
		var v s
		err := ToStruct(`{"a": true}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), "无法转换 true 为 uint")
	})
}

func TestToStructFloat(t *testing.T) {
	type s struct {
		A float32  `json:"a"`
		B float64  `json:"b"`
		C float64  `json:"c"`
		P *float64 `json:"p"`
	}

	var v s
	err := ToStruct(`{"a": "1.5", "b": 2.25, "c": "3.75", "p": 4.5}`, &v)
	require.NoError(t, err)
	require.Equal(t, float32(1.5), v.A)
	require.Equal(t, 2.25, v.B)
	require.Equal(t, 3.75, v.C)
	require.NotNil(t, v.P)
	require.Equal(t, 4.5, *v.P)
}

func TestToStructFloatError(t *testing.T) {
	t.Run("字符串解析失败", func(t *testing.T) {
		type s struct {
			A float64 `json:"a"`
		}
		var v s
		err := ToStruct(`{"a": "abc"}`, &v)
		require.ErrorIs(t, err, ErrSyntax)
	})

	t.Run("不支持的数据类型", func(t *testing.T) {
		type s struct {
			A float64 `json:"a"`
		}
		var v s
		err := ToStruct(`{"a": true}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), "无法转换 true 为 float")
	})
}
