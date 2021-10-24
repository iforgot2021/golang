package fab

import (
	"testing"
)

type CaseFab struct {
	Input  int
	Expect int
}

func TestFab(t *testing.T) {
	tc := []CaseFab{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
	}

	for _, v := range tc {
		if Fab(v.Input) != v.Expect {
			t.Errorf("expect %d, got %d\n", v.Expect, Fab(v.Input))
		}
		t.Log(v)
	}
}

type CaseList struct {
	Input  int
	Expect []int
}

func TestFabList(t *testing.T) {
	tc := []CaseList{
		{1, []int{1}},
		{3, []int{1, 1, 2}},
		{5, []int{1, 1, 2, 3, 5}},
		{7, []int{1, 1, 2, 3, 5, 8, 13}},
	}

	for _, v := range tc {
		res := FabList(v.Input)
		for i, rv := range res {
			if rv != v.Expect[i] {
				t.Errorf("expect %d got %d\n", v.Expect[i], rv)
			}
		}
	}
}
