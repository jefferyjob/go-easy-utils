package jsonUtil

import (
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
}

func TestJsonToStructMap2(t *testing.T) {
	var exampleVal struct {
		MapInt    map[string]int    `json:"mapInt"`
		MapString map[string]string `json:"mapString"`
		MapBool   map[string]bool   `json:"mapBool"`
		MapStruct map[string]struct {
			S1 string `json:"s1"`
			S2 int64  `json:"s2"`
		} `json:"mapStruct"`
		MapMap map[string]map[string]int32 `json:"mapMap"`
		Uid    int                         `json:"uid"`
	}

	exampleJson := `
	{
	  "mapInt": {
		"key1": 1,
		"key2": 2
	  },
	  "mapString": {
		"key1": "value1",
		"key2": "value2"
	  },
	  "mapBool": {
		"key1": true,
		"key2": false
	  },
	  "mapStruct": {
		"key1": {
		  "s1": "string_value1",
		  "s2": 12345
		},
		"key2": {
		  "s1": "string_value2",
		  "s2": 67890
		}
	  },
	  "mapMap": {
		"key1": {
		  "subkey1": 10,
		  "subkey2": 20
		},
		"key2": {
		  "subkey1": 30,
		  "subkey2": 40
		}
	  },
	  "uid": 123
	}`

	err := JsonToStruct(exampleJson, &exampleVal)
	if err != nil {
		t.Errorf("err %s", err)
		return
	}
}
