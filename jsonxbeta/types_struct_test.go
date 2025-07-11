package jsonxbeta

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

const tsJ1 = `
{
	"name": "Alice",
	"age": "30"
}
`

type tsS1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestStruct1(t *testing.T) {
	var s tsS1
	err := ToStruct(tsJ1, &s)
	require.NoError(t, err)

	fmt.Printf("%+v\n", s)
}
