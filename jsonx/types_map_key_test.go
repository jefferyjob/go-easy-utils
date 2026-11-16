package jsonx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type mapKeyStr string

func TestToStructMapKeyTypes(t *testing.T) {
	type s struct {
		U8    map[uint8]int             `json:"u8"`
		U64   map[uint64]string         `json:"u64"`
		I16   map[int16]float64         `json:"i16"`
		F32   map[float32]string        `json:"f32"`
		F64   map[float64]bool          `json:"f64"`
		Str   map[mapKeyStr]int         `json:"str"`
		Nest  map[string]map[string]int `json:"nest"`
		Slice map[string][]int          `json:"slice"`
	}

	jsonStr := `{
		"u8": {"255": 1, "0": 2},
		"u64": {"18446744073709551615": "max"},
		"i16": {"-300": 1.5},
		"f32": {"2.25": "x"},
		"f64": {"1.5": true, "2": false},
		"str": {"k": 3},
		"nest": {"a": {"b": 1}},
		"slice": {"n": [1, 2]}
	}`

	var v s
	err := ToStruct(jsonStr, &v)
	require.NoError(t, err)
	require.Equal(t, map[uint8]int{255: 1, 0: 2}, v.U8)
	require.Equal(t, map[uint64]string{18446744073709551615: "max"}, v.U64)
	require.Equal(t, map[int16]float64{-300: 1.5}, v.I16)
	require.Equal(t, map[float32]string{2.25: "x"}, v.F32)
	require.Equal(t, map[float64]bool{1.5: true, 2: false}, v.F64)
	require.Equal(t, map[mapKeyStr]int{"k": 3}, v.Str)
	require.Equal(t, map[string]map[string]int{"a": {"b": 1}}, v.Nest)
	require.Equal(t, map[string][]int{"n": {1, 2}}, v.Slice)
}

func TestToStructMapError(t *testing.T) {
	t.Run("值不是map", func(t *testing.T) {
		type s struct {
			M map[string]int `json:"m"`
		}
		var v s
		err := ToStruct(`{"m": "not-map"}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), "不是 map 类型")
	})

	t.Run("int类型key解析失败", func(t *testing.T) {
		type s struct {
			M map[int]string `json:"m"`
		}
		var v s
		err := ToStruct(`{"m": {"abc": "v"}}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), `无法将 "abc" 转为 int`)
	})

	t.Run("int8类型key溢出", func(t *testing.T) {
		type s struct {
			M map[int8]string `json:"m"`
		}
		var v s
		err := ToStruct(`{"m": {"300": "v"}}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), `无法将 "300" 转为 int8`)
	})

	t.Run("bool类型key解析失败", func(t *testing.T) {
		type s struct {
			M map[bool]string `json:"m"`
		}
		var v s
		err := ToStruct(`{"m": {"yes": "v"}}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), `无法将 "yes" 转为 bool`)
	})

	t.Run("uint类型key为负数", func(t *testing.T) {
		type s struct {
			M map[uint]string `json:"m"`
		}
		var v s
		err := ToStruct(`{"m": {"-1": "v"}}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), `无法将 "-1" 转为 uint`)
	})

	t.Run("float类型key解析失败", func(t *testing.T) {
		type s struct {
			M map[float64]string `json:"m"`
		}
		var v s
		err := ToStruct(`{"m": {"x": "v"}}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), `无法将 "x" 转为 float64`)
	})

	t.Run("不支持的key类型", func(t *testing.T) {
		type s struct {
			M map[complex128]int `json:"m"`
		}
		var v s
		err := ToStruct(`{"m": {"k": 1}}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), "不支持的 map key 类型 complex128")
	})

	t.Run("map元素值解析失败", func(t *testing.T) {
		type s struct {
			M map[string]int `json:"m"`
		}
		var v s
		err := ToStruct(`{"m": {"a": "x"}}`, &v)
		require.ErrorIs(t, err, ErrSyntax)
		require.Contains(t, err.Error(), "m[a]")
	})
}
