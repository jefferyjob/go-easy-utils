package jsonx

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

const tmJ1 = `
{
  "mapNum": {"1": "one", "42": "forty-two"},
  "mapBool": {"true": 1.23, "false": 4.56}
}
`

type tmS1 struct {
	MapNum  map[int]string   `json:"mapNum"`
	MapBool map[bool]float64 `json:"mapBool"`
}

func TestMap1Beta(t *testing.T) {
	var s tmS1
	err := ToStruct(tmJ1, &s)
	require.NoError(t, err)

	fmt.Printf("%+v\n", s)
}

func TestToStructMapPointerValues(t *testing.T) {
	type s struct {
		Nums map[string]*int    `json:"nums"`
		M    map[string]*string `json:"m"`
	}

	jsonStr := `{"nums": {"a": "1", "b": 2, "c": null}, "m": {"k": "v", "e": null}}`

	var v s
	err := ToStruct(jsonStr, &v)
	require.NoError(t, err)
	require.Len(t, v.Nums, 3)
	require.NotNil(t, v.Nums["a"])
	require.Equal(t, 1, *v.Nums["a"])
	require.NotNil(t, v.Nums["b"])
	require.Equal(t, 2, *v.Nums["b"])
	// null 值保持 nil 指针
	require.Nil(t, v.Nums["c"])

	require.Len(t, v.M, 2)
	require.Equal(t, "v", *v.M["k"])
	require.Nil(t, v.M["e"])
}

func TestToStructMapValueError(t *testing.T) {
	t.Run("值溢出", func(t *testing.T) {
		type s struct {
			Nums map[string]int8 `json:"nums"`
		}
		var v s
		err := ToStruct(`{"nums": {"k": 300}}`, &v)
		require.ErrorIs(t, err, ErrValOut)
		require.Contains(t, err.Error(), "nums[k]")
	})

	t.Run("值不是map", func(t *testing.T) {
		type s struct {
			Nums map[string]int `json:"nums"`
		}
		var v s
		err := ToStruct(`{"nums": "not-map"}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nums: 不是 map 类型")
	})
}
