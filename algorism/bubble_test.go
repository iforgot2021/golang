package algorism

import "testing"

type testCase struct {
	Input, Expect []int64
}

func TestBubbleSort(t *testing.T) {
	testCases := []testCase{
		{
			[]int64{5, 4, 3, 2, 1},
			[]int64{1, 2, 3, 4, 5},
		},
		{
			[]int64{9, -1, 5, -12, 43, 23},
			[]int64{-12, -1, 5, 9, 23, 43},
		},
	}

	for _, cs := range testCases {
		got := BubbleSort(cs.Input)
		for i, v := range got {
			if cs.Expect[i] != v {
				t.Errorf("error expect:%d got:%d", cs.Expect[i], v)
			}
		}
	}
}
