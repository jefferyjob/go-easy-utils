package jsonUtil

import (
	"testing"
)

func TestJsonToStructSlice1(t *testing.T) {
	data := `
		{
			"uid": 43015653,
			"fids": ["43015653", 43015666]
		}
	`
	var target struct {
		Uid  int      `json:"uid"`
		Fids []string `json:"fids"`
	}

	err := JsonToStruct(data, &target)
	if err != nil {
		t.Errorf("err %s", err)
		return
	}
}
