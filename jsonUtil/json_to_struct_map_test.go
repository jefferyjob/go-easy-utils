package jsonUtil

import (
	"fmt"
	"testing"
)

func TestJsonToStructMap1(t *testing.T) {
	data := `
	{
		"uid": 43015653,
		"foll": {
			"43015653": true,
			"43015666": false
		},
		"followed": {
			"friendRed": {
				"43015653": true,
				"43015666": false
			},
			"friendWhite": {
				"43015653": true,
				"43015666": false
			}
		}
	}
	`
	var target struct {
		Uid      int                        `json:"uid"`
		Foll     map[string]bool            `json:"foll"`
		Followed map[string]map[string]bool `json:"followed"`
	}

	err := JsonToStruct(data, &target)
	if err != nil {
		t.Errorf("err %s", err)
		return
	}
	fmt.Printf("%+v, err:%s \n", target, err)
}

func TestJsonToStructMap2(t *testing.T) {
	data := `
	{
		"uid": 43015653,
		"foll": {
			"boy": {
				"t1": "v1",
				"t2": "v2"
			},
			"girl": {
				"t1": "v1",
				"t2": "v2"
			}
		}
	}
	`
	type Val struct {
		T1 string `json:"t1"`
		T2 string `json:"t2"`
	}

	var target struct {
		Uid  int            `json:"uid"`
		Foll map[string]Val `json:"foll"`
	}

	err := JsonToStruct(data, &target)
	if err != nil {
		t.Errorf("err %s", err)
		return
	}
	fmt.Printf("%+v, err:%s \n", target, err)
}


