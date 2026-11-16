package jsonx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToStructSlice(t *testing.T) {
	type friend struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type s struct {
		Mixed   []string `json:"mixed"`
		Nums    []int    `json:"nums"`
		Empty   []int    `json:"empty"`
		Nested  [][]int  `json:"nested"`
		Friends []friend `json:"friends"`
		Any     []any    `json:"any"`
	}

	jsonStr := `{
		"mixed": ["a", 1, true],
		"nums": ["10", 20],
		"empty": [],
		"nested": [[1, 2], [3]],
		"friends": [{"name": "Alice", "age": "30"}],
		"any": [1, "x", null, [2]]
	}`

	var v s
	err := ToStruct(jsonStr, &v)
	require.NoError(t, err)
	require.Equal(t, []string{"a", "1", "true"}, v.Mixed)
	require.Equal(t, []int{10, 20}, v.Nums)
	require.Equal(t, []int{}, v.Empty)
	require.Equal(t, [][]int{{1, 2}, {3}}, v.Nested)
	require.Equal(t, []friend{{Name: "Alice", Age: 30}}, v.Friends)
	require.Equal(t, []any{float64(1), "x", nil, []any{float64(2)}}, v.Any)
}

func TestToStructSliceError(t *testing.T) {
	t.Run("值不是数组", func(t *testing.T) {
		type s struct {
			Nums []int `json:"nums"`
		}
		var v s
		err := ToStruct(`{"nums": "not-array"}`, &v)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nums: 不是数组类型")
	})

	t.Run("数组元素解析失败", func(t *testing.T) {
		type s struct {
			Nums []int `json:"nums"`
		}
		var v s
		err := ToStruct(`{"nums": ["1", "abc"]}`, &v)
		require.ErrorIs(t, err, ErrSyntax)
		require.Contains(t, err.Error(), "nums[1]")
	})

	t.Run("嵌套数组元素解析失败", func(t *testing.T) {
		type s struct {
			Nested [][]int `json:"nested"`
		}
		var v s
		err := ToStruct(`{"nested": [[1], ["x"]]}`, &v)
		require.ErrorIs(t, err, ErrSyntax)
		require.Contains(t, err.Error(), "nested[1][0]")
	})

	t.Run("结构体数组字段解析失败", func(t *testing.T) {
		type s struct {
			Friends []Address `json:"friends"`
		}
		var v s
		err := ToStruct(`{"friends": [{"postcode": "abc"}]}`, &v)
		require.ErrorIs(t, err, ErrSyntax)
		require.Contains(t, err.Error(), "friends[0].postcode")
	})

	t.Run("数组元素溢出", func(t *testing.T) {
		type s struct {
			Nums []int8 `json:"nums"`
		}
		var v s
		err := ToStruct(`{"nums": [1, 300]}`, &v)
		require.ErrorIs(t, err, ErrValOut)
		require.Contains(t, err.Error(), "nums[1]")
	})
}

func TestToStructSlicePointerElements(t *testing.T) {
	type friend struct {
		Name string `json:"name"`
	}

	type s struct {
		Nums    []*int     `json:"nums"`
		Names   []*string  `json:"names"`
		Friends []*friend  `json:"friends"`
	}

	jsonStr := `{
		"nums": ["7", 8, null],
		"names": ["Alice", null],
		"friends": [{"name": "Bob"}, null]
	}`

	var v s
	err := ToStruct(jsonStr, &v)
	require.NoError(t, err)
	require.Len(t, v.Nums, 3)
	require.NotNil(t, v.Nums[0])
	require.Equal(t, 7, *v.Nums[0])
	require.NotNil(t, v.Nums[1])
	require.Equal(t, 8, *v.Nums[1])
	// null 元素保持 nil 指针
	require.Nil(t, v.Nums[2])

	require.Len(t, v.Names, 2)
	require.Equal(t, "Alice", *v.Names[0])
	require.Nil(t, v.Names[1])

	require.Len(t, v.Friends, 2)
	require.Equal(t, "Bob", v.Friends[0].Name)
	require.Nil(t, v.Friends[1])
}
