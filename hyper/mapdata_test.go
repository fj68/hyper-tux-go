package hyper_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go/hyper"
)

func TestNewMapdataFromSlice(t *testing.T) {
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
				m := hyper.NewMapdata(hyper.Size{8, 8})
				m.HWalls[1] = append(m.HWalls[1], 1, 6)
				m.HWalls[5] = append(m.HWalls[5], 6)
				m.VWalls[1] = append(m.VWalls[1], 1)
				m.VWalls[2] = append(m.VWalls[2], 6)
				m.VWalls[6] = append(m.VWalls[6], 5)
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
