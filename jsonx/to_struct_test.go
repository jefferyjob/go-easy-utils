package jsonx

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

const Json1 = `{
	"name": "Bob",
	"age": "25",
	"height": 175.5,
	"weight": "60.2",
	"isBoy": true,
	"isStudent": "false",
	"nickname": null,
	"hobbies": ["reading", 123, true, null],
	"likeNum": [101, 202],
	"scores": ["99.5", 88.1, "76"],
	"tags": [],
	"address": {
		"country": "China",
		"city": "Shanghai",
		"postcode": 200000
	},
	"contacts": {
		"email": "bob@example.com",
		"phone": 1234567890,
		"wechat": null
	},
	"preferences": {
		"food": ["noodle", "pizza"],
		"sports": {
			"indoor": ["chess", "table tennis"],
			"outdoor": ["basketball", "football"]
		}
	},
	"meta": {
		"lastLogin": "2025-07-08T10:30:00Z",
		"active": "true",
		"loginCount": 10
	},
	"notes": null,
	"custom": {
		"field1": "value1",
		"field2": 2
	}
}`

type S1 struct {
	Name        *string           `json:"name"`
	Age         int               `json:"age"`
	Height      float64           `json:"height"`
	Weight      float64           `json:"weight"`
	IsBoy       bool              `json:"isBoy"`
	IsStudent   bool              `json:"isStudent"`
	Nickname    string            `json:"nickname"` // null 会被忽略
	Hobbies     []string          `json:"hobbies"`  // 数组中混合类型，将自动转换为字符串
	LikeNum     []int64           `json:"likeNum"`
	Scores      []float64         `json:"scores"`
	Tags        []string          `json:"tags"`
	Address     *Address          `json:"address"`
	Contacts    map[string]string `json:"contacts"`
	Preferences struct {
		Food   []string            `json:"food"`
		Sports map[string][]string `json:"sports"` // indoor / outdoor
	} `json:"preferences"` // 嵌套结构体
	Meta struct {
		LastLogin  string `json:"lastLogin"`
		Active     bool   `json:"active"`
		LoginCount int    `json:"loginCount"`
	} `json:"meta"` // 键值混合类型
	Notes  string         `json:"notes"`  // null 会被忽略
	Custom map[string]any `json:"custom"` // 深层次 map 类型
}

type Address struct {
	Country  string `json:"country"`
	City     string `json:"city"`
	Postcode int    `json:"postcode"`
}

func TestToStruct1Beta(t *testing.T) {
	s := &S1{}
	err := ToStruct(Json1, s)
	require.NoError(t, err)
	fmt.Printf("%+v\n", s)

	fmt.Println("name: ", s.Name)
	fmt.Println("country: ", s.Address.Country)
}

const Json2 = `{
	"name": "Bob",
	"address": {
		"country": "China",
		"city": "Shanghai",
		"postcode": 200000
	}
}`

type S2 struct {
	Name    *string  `json:"name"`
	Address *Address `json:"address"`
}

func TestA1(t *testing.T) {
	s := &S2{}
	err := json.Unmarshal([]byte(Json2), s)
	require.NoError(t, err)

	fmt.Printf("%+v\n", s)
	fmt.Println("name: ", s.Name)
	fmt.Println("country: ", s.Address.Country)
}
