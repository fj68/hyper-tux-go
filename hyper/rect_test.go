package hyper_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go/hyper"
)

func TestRect_Contains(t *testing.T) {
	testcases := []struct {
		Name string
		hyper.Rect
		hyper.Point
		Expected bool
	}{
		{
			"totally contained",
			hyper.Rect{hyper.Point{0, 0}, hyper.Point{5, 5}},
			hyper.Point{3, 3},
			true,
		},
		{
			"on topleft edge",
			hyper.Rect{hyper.Point{0, 0}, hyper.Point{5, 5}},
			hyper.Point{0, 0},
			true,
		},
		{
			"on bottomright edge",
			hyper.Rect{hyper.Point{0, 0}, hyper.Point{5, 5}},
			hyper.Point{5, 5},
			false,
		},
		{
			"totally uncontained",
			hyper.Rect{hyper.Point{0, 0}, hyper.Point{5, 5}},
			hyper.Point{6, 6},
			false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			if testcase.Rect.Contains(testcase.Point) != testcase.Expected {
				t.Errorf("unexpected value: Expected = %t, Rect = %+v, Point = %+v", testcase.Expected, testcase.Rect, testcase.Point)
			}
		})
	}
}
