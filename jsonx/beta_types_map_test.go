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
	err := ToStructBeta(tmJ1, &s)
	require.NoError(t, err)

	fmt.Printf("%+v\n", s)
}
