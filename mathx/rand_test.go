package mathx

import "testing"

func TestRand(t *testing.T) {
	testCases := []struct {
		name   string
		minNum int
		maxNum int
	}{
		{
			name:   "测试随机数在范围内",
			minNum: 1,
			maxNum: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				num := Rand(tc.minNum, tc.maxNum)
				if num < tc.minNum || num > tc.maxNum {
					t.Errorf("Rand(%d, %d) returned %d, which is outside the range of [%d, %d]", tc.minNum, tc.maxNum, num, tc.minNum, tc.maxNum)
				}
			}
		})
	}
}
