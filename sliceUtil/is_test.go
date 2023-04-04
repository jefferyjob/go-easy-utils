package sliceUtil

import "testing"

func TestIsSlice(t *testing.T) {
	var tests = []struct {
		input interface{}
		want  bool
	}{
		{[]int{1, 1, 3}, true},
		{[]interface{}{1, 2, "a"}, true},
		{[]map[string]interface{}{
			{"1": 1},
			{"c": 89},
		}, true},
		{"1234", false},
		{make(chan int), false},
	}
	for _, test := range tests {
		if got := Is(test.input); got != test.want {
			t.Errorf("IsSlice(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}
