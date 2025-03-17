package hyper_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go/hyper"
)

func TestNewMapdateFromSlice(t *testing.T) {
	testcases := []struct {
		Name     string
		Input    [][]int
		Expected func() *hyper.Mapdata
	}{
		{
			Name: "small mapdata",
			Input: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 3, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 2, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 3, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			Expected: func() *hyper.Mapdata {
				m := hyper.NewMapdata(hyper.Size{7, 7})
				m.HWalls[1].Add(1)
				m.HWalls[1].Add(5)
				m.HWalls[5].Add(5)
				m.VWalls[1].Add(1)
				m.HWalls[5].Add(5)
				return m
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			actual, err := hyper.NewMapdataFromSlice(testcase.Input)
			if err != nil {
				t.Error(err)
			}
			expected := testcase.Expected()
			if !expected.Equals(actual) {
				t.Errorf("expected:\n\t%+v\nactual:\n\t%+v\n", expected, actual)
			}
		})
	}
}
